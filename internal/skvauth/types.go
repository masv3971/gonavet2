package skvauth

import (
	"fmt"
	"sync"
)

// JWTRequest request type for ensureJWT
//type JWTRequest string

type JWTReply struct {
	AccessToken string `json:"access_token" validate:"required"`
	ExpiresIn   int64  `json:"expires_in" validate:"required"`
	TokenType   string `json:"token_type" validate:"required"`
	Scope       string `json:"scope" validate:"required"`
}

// JWT holds the raw token, mutex lock and expiration dates
type JWT struct {
	sync.RWMutex
	RAW       string
	ExpiresAt int64
	IssuedAt  int64
	NotBefore int64
}

// Errors is a general error reply
type Errors struct {
	Detail []struct {
		Loc  []string `json:"loc"`
		Msg  string   `json:"msg"`
		Type string   `json:"type"`
	} `json:"detail"`
}

// Error interface
type Error interface {
	Error() string
}

func (e *Errors) Error() string {
	return fmt.Sprintf("error: %v", e.Detail)
}
