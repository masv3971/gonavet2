package gonavet2

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

// Config configs New
type Config struct {
	URL            string            `validate:"required"`
	Certificate    *x509.Certificate `validate:"required"`
	CertificatePEM []byte            `validate:"required"`
	PrivateKey     *rsa.PrivateKey   `validate:"required"`
	PrivateKeyPEM  []byte            `validate:"required"`
	ProxyURL       string
}

// Client holds navet object
type Client struct {
	httpClient     *http.Client
	rateLimit      *rate.Limiter
	format         string
	url            string
	certificate    *x509.Certificate
	certificatePEM []byte
	chain          *x509.CertPool
	chainPEM       []byte
	privateKey     *rsa.PrivateKey
	privateKeyPEM  []byte
	proxyURL       string

	Fetch  *fetchService
	Search *searchService
}

// New creates a new instance of gonavet
func New(config Config) (*Client, error) {
	c := &Client{
		rateLimit: rate.NewLimiter(rate.Every(1*time.Second), 15),
	}

	c.Fetch = &fetchService{client: c}
	c.Search = &searchService{client: c}

	return c, nil
}

func (c *Client) httpConfigure() error {
	keyPair, err := tls.X509KeyPair(c.certificatePEM, c.privateKeyPEM)
	if err != nil {
		return err
	}

	tlsCfg := &tls.Config{
		Rand:         rand.Reader,
		Certificates: []tls.Certificate{keyPair},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		//ClientCAs:          c.chainDER,
		InsecureSkipVerify: false,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
		PreferServerCipherSuites:    true,
		SessionTicketsDisabled:      false,
		DynamicRecordSizingDisabled: false,
		Renegotiation:               0,
		KeyLogWriter:                nil,
	}

	tlsCfg.BuildNameToCertificate()

	c.httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     tlsCfg,
			DialContext:         nil,
			TLSHandshakeTimeout: 30 * time.Second,
			Proxy:               http.ProxyFromEnvironment,
		},
	}

	return nil
}

// NewRequest make a new request
func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}
	url := u.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		payload := struct {
			Data interface{} `json:"data"`
		}{
			Data: body,
		}
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "gonavet") // TODO(masv): add version
	return req, nil
}

// Do does the new request
func (c *Client) do(ctx context.Context, req *http.Request, value interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		errorReply := &Errors{}
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(buf.Bytes(), errorReply); err != nil {
			return nil, err
		}
		return nil, errorReply
	}

	if err := json.NewDecoder(resp.Body).Decode(value); err != nil {
		return nil, err
	}

	return resp, nil
}

func checkResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 500:
		return oneError("Not allowed", "statusCode", "checkResponse", "")
	}
	return oneError("Invalid request", "statusCode", "checkResponse", "")
}

func (c *Client) call(ctx context.Context, acceptHeader, method, url string, body, reply interface{}) (*http.Response, error) {
	request, err := c.newRequest(
		ctx,
		method,
		url,
		body,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(ctx, request, reply)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
