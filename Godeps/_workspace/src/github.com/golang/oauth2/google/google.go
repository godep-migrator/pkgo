// Copyright 2014 The oauth2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package google provides support for making
// OAuth2 authorized and authenticated HTTP requests
// to Google APIs. It supports Web server, client-side,
// service accounts, Google Compute Engine service accounts,
// and Google App Engine service accounts authorization
// and authentications flows:
//
// For more information, please read
// https://developers.google.com/accounts/docs/OAuth2.
package google

import (
	"encoding/json"
	"net/http"
	"path"
	"time"

	"github.com/golang/oauth2"
)

const (
	// Google endpoints.
	uriGoogleAuth  = "https://accounts.google.com/o/oauth2/auth"
	uriGoogleToken = "https://accounts.google.com/o/oauth2/token"
)

type metaTokenRespBody struct {
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Duration `json:"expires_in"`
	TokenType   string        `json:"token_type"`
}

// ComputeEngineConfig represents a OAuth 2.0 consumer client
// running on Google Compute Engine.
type ComputeEngineConfig struct {
	// Client is the HTTP client to be used to retrieve
	// tokens from the OAuth 2.0 provider.
	Client *http.Client

	// Transport is the round tripper to be used
	// to construct new oauth2.Transport instances from
	// this configuration.
	Transport http.RoundTripper

	account string
}

// NewConfig creates a new OAuth2 config that uses Google
// endpoints.
func NewConfig(opts *oauth2.Options) (*oauth2.Config, error) {
	return oauth2.NewConfig(opts, uriGoogleAuth, uriGoogleToken)
}

// NewServiceAccountConfig creates a new JWT config that can
// fetch Bearer JWT tokens from Google endpoints.
func NewServiceAccountConfig(opts *oauth2.JWTOptions) (*oauth2.JWTConfig, error) {
	return oauth2.NewJWTConfig(opts, uriGoogleToken)
}

// NewComputeEngineConfig creates a new config that can fetch tokens
// from Google Compute Engine instance's metaserver. If no account is
// provided, default is used.
func NewComputeEngineConfig(account string) *ComputeEngineConfig {
	return &ComputeEngineConfig{
		Client:    http.DefaultClient,
		Transport: http.DefaultTransport,
		account:   account,
	}
}

// NewTransport creates an authorized transport.
func (c *ComputeEngineConfig) NewTransport() *oauth2.Transport {
	return oauth2.NewTransport(c.Transport, c, nil)
}

// FetchToken retrieves a new access token via metadata server.
func (c *ComputeEngineConfig) FetchToken(existing *oauth2.Token) (token *oauth2.Token, err error) {
	account := "default"
	if c.account != "" {
		account = c.account
	}
	u := "http://" + path.Join("metadata/computeMetadata/v1/instance/service-accounts", account, "token")
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return
	}
	req.Header.Add("X-Google-Metadata-Request", "True")
	resp, err := c.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var tokenResp metaTokenRespBody
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		return
	}
	token = &oauth2.Token{
		AccessToken: tokenResp.AccessToken,
		TokenType:   tokenResp.TokenType,
		Expiry:      time.Now().Add(tokenResp.ExpiresIn * time.Second),
	}
	return
}
