package skvauth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/go-resty/resty/v2"
	"golang.org/x/time/rate"
)

// Client holds the object of skatteverket
type Client struct {
	httpClient  *resty.Client
	rateLimit   *rate.Limiter
	certificate []byte
	privateKey  []byte
	baseURL     string
	proxyURL    string

	scope        string
	grantType    string
	clientSecret string
	ClientID     string
	JWT          *JWT
}

type Config struct {
	CertificatePEM []byte
	PrivateKeyPEM  []byte
	BaseURL        string
	ProxyURL       string
	Scope          string
	ClientSecret   string
	ClientID       string
	GrantType      string
}

// New creates a new instance of skatteverket
func New(config *Config) *Client {
	c := &Client{
		httpClient:   resty.New(),
		rateLimit:    rate.NewLimiter(rate.Every(1*time.Second), 15),
		certificate:  config.CertificatePEM,
		privateKey:   config.PrivateKeyPEM,
		baseURL:      config.BaseURL,
		proxyURL:     config.ProxyURL,
		scope:        config.Scope,
		grantType:    config.GrantType,
		clientSecret: config.ClientSecret,
		ClientID:     config.ClientID,
		JWT:          &JWT{},
	}

	c.httpClient.SetBaseURL(c.baseURL)
	c.httpClient.SetProxy(c.proxyURL)
	c.httpClient.SetRetryCount(3)
	c.httpClient.SetRetryWaitTime(3 * time.Second)
	c.httpClient.SetRetryMaxWaitTime(20 * time.Second)
	c.httpClient.SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
		return 0, errors.New("quota exceeded")
	})

	return c
}

// newJWT gets a jwt
func (c *Client) newJWT(ctx context.Context) error {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
		"Accept":       "application/json",
		"User-Agent":   "gonavet2",
	}

	result := &JWTReply{}
	body := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&scope=%s",
		c.ClientID,
		c.clientSecret,
		c.scope,
	)

	if err := c.rateLimit.Wait(ctx); err != nil {
		return nil
	}

	c.JWT.Lock()
	defer c.JWT.Unlock()

	_, err := c.httpClient.R().SetContext(ctx).SetHeaders(headers).SetBody(body).SetResult(result).Post("oauth2/v1/sysorg/token")
	if err != nil {
		return err
	}

	if err := validate(result); err != nil {
		return err
	}

	if err := c.parseJWT(result.AccessToken); err != nil {
		return err
	}

	return nil
}

func (c *Client) parseJWT(accessToken string) error {
	jwtBase64 := strings.Split(accessToken, ".")
	if len(jwtBase64) < 3 {
		return errors.New("ERROR invalid JWT, corrupt split")
	}

	jwtDecoded, err := base64.RawURLEncoding.DecodeString(jwtBase64[1])
	if err != nil {
		return err
	}
	type jwtSpec struct {
		Exp             int64  `json:"exp"`
		Iat             int64  `json:"iat"`
		Iss             string `json:"iss"`
		Nbf             int64  `json:"nbf"`
		RequestedAccess []struct {
			Scope string `json:"scope"`
			Type  string `json:"type"`
		} `json:"requested_access"`
		Scopes  []string `json:"scopes"`
		Source  string   `json:"source"`
		Sub     string   `json:"sub"`
		Version int      `json:"version"`
	}

	jwtSpecS := &jwtSpec{}
	if err := json.Unmarshal(jwtDecoded, jwtSpecS); err != nil {
		return err
	}

	jwt := &JWT{
		RAW:       accessToken,
		ExpiresAt: jwtSpecS.Exp,
		IssuedAt:  jwtSpecS.Iat,
		NotBefore: jwtSpecS.Nbf,
	}
	c.JWT = jwt

	return nil
}

// Valid checks if jwt token i valid of not
func (jwt *JWT) Valid() bool {
	jwt.RLock()
	defer jwt.RUnlock()

	// token has expired
	unixNow := time.Now().Unix()
	if jwt.ExpiresAt < unixNow {
		return false
	}

	// token issued in the future
	if jwt.IssuedAt > unixNow {
		return false
	}

	// token starts to be valid in the future
	if jwt.NotBefore > unixNow {
		return false
	}

	// Token has too little time left
	if jwt.ExpiresAt < unixNow+360 {
		return false
	}
	return true
}

// EnsureJWT ensure that a jwt is present and valid
func (c *Client) EnsureJWT(ctx context.Context) error {
	if c.JWT != nil {
		if c.JWT.Valid() {
			return nil
		}
	}

	if err := c.newJWT(ctx); err != nil {
		return err
	}
	return nil
}

func validate(s interface{}) error {
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("ERR: Field %q of type %q violates rule: %q\n", err.Namespace(), err.Kind(), err.Tag())
		}
		return errors.New("Validation error")
	}
	return nil
}
