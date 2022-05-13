package gonavet2

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/masv3971/gonavet2/internal/skvauth"
	"github.com/mcuadros/go-defaults"
	"golang.org/x/time/rate"
)

// Config configs New
type Config struct {
	Debug             bool
	BaseURL           string `validate:"required"`
	ProxyURL          string
	CertificatePEM    []byte      `validate:"required"`
	PrivateKeyPEM     []byte      `validate:"required"`
	OrgNumber         string      `validate:"required"`
	OrderID           string      `validate:"required"`
	Token             OauthConfig `validate:"required"`
	APIGWClientID     string      `validate:"required"`
	APIGWClientSecret string      `validate:"required"`
}

type OauthConfig struct {
	Scope        string `default:"fbfuppgoffakt"`
	BaseURL      string
	ClientID     string
	ClientSecret string
	GrantType    string `default:"client_credentials"`
}

// Client holds navet object
type Client struct {
	config         *Config
	httpClient     *resty.Client
	rateLimit      *rate.Limiter
	baseURL        string
	chainPEM       []byte
	certificatePEM []byte
	privateKeyPEM  []byte
	proxyURL       string
	skvOauth       *skvauth.Client

	// organisationsnummer
	orgNumber string

	// bestallningsidentitet, id from skv
	orderID string

	Fetch  *fetchService
	Search *searchService
}

// New creates a new instance of gonavet
func New(config *Config) (*Client, error) {
	c := &Client{
		config:     config,
		rateLimit:  rate.NewLimiter(rate.Every(1*time.Second), 15),
		httpClient: resty.New(),
	}

	c.httpClient.SetBaseURL(c.baseURL)
	c.httpClient.SetProxy(c.proxyURL)
	c.httpClient.SetRetryCount(3)
	c.httpClient.SetRetryWaitTime(3 * time.Second)
	c.httpClient.SetRetryMaxWaitTime(20 * time.Second)
	c.httpClient.SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
		return 0, errors.New("quota exceeded")
	})

	defaults.SetDefaults(config)

	c.httpClient.SetDebug(c.config.Debug)

	keyPair, err := tls.X509KeyPair(c.certificatePEM, c.privateKeyPEM)
	if err != nil {
		return nil, err
	}
	c.httpClient.SetCertificates(keyPair)

	c.skvOauth = skvauth.New(&skvauth.Config{
		CertificatePEM: c.config.CertificatePEM,
		PrivateKeyPEM:  c.config.PrivateKeyPEM,
		BaseURL:        c.config.Token.BaseURL,
		ProxyURL:       c.config.ProxyURL,
		Scope:          c.config.Token.Scope,
		ClientSecret:   c.config.Token.ClientSecret,
		ClientID:       c.config.Token.ClientID,
		GrantType:      c.config.Token.GrantType,
	})

	c.Fetch = &fetchService{client: c, endpoint: "/hamta"}
	c.Search = &searchService{client: c, endpoint: "/sok"}

	return c, nil
}

func (c *Client) req(ctx context.Context, body, result interface{}) (*resty.Request, error) {
	ctxv := context.WithValue(ctx, "skv-id", skvClientCorrelationID(uuid.NewString()))

	c.skvOauth.JWT.RLock()
	defer c.skvOauth.JWT.RUnlock()

	if err := c.skvOauth.EnsureJWT(ctx); err != nil {
		return nil, err
	}

	if err := c.rateLimit.Wait(ctx); err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Accept":                    "application/json",
		"Content-Type":              "application/json",
		"User-Agent":                "gonavet2",
		"skv_client_correlation_id": ctxv.Value("skv-id").(skvClientCorrelationID).string(),
		"client_id":                 c.config.APIGWClientID,
		"client_secret":             c.config.APIGWClientSecret,
	}

	req := c.httpClient.
		R().
		SetContext(ctxv).
		SetHeaders(headers).
		SetAuthToken(c.skvOauth.JWT.RAW).
		SetBody(body).SetResult(result)

	return req, nil
}

func checkResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 500:
		//	return oneError("Not allowed", "statusCode", "checkResponse", "")
	}
	//return oneError("Invalid request", "statusCode", "checkResponse", "")
	return nil
}

type skvClientCorrelationID string

func (s skvClientCorrelationID) string() string { return string(s) }
