// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package httpjson

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const errStr = "error decoding JWK component: %w"

// customTokenSource is a custom implementation of the oauth2.TokenSource interface.
type customTokenSource struct {
	mu      sync.Mutex
	ctx     context.Context
	conf    *oauth2.Config
	token   *oauth2.Token
	oktaJWK string
}

// fetchOktaOauthClient fetches an OAuth2 client using the Okta JWK credentials.
func (o *oAuth2Config) fetchOktaOauthClient(ctx context.Context, _ *http.Client) (*http.Client, error) {
	oktaJWK := string(o.OktaJWKJSON)
	conf := &oauth2.Config{
		ClientID: o.ClientID,
		Scopes:   o.Scopes,
		Endpoint: oauth2.Endpoint{
			TokenURL: o.TokenURL,
		},
	}

	oktaJWT, err := generateOktaJWT(oktaJWK, conf)
	if err != nil {
		return nil, fmt.Errorf("oauth2 client: error generating Okta JWT: %w", err)
	}

	token, err := exchangeForBearerToken(ctx, oktaJWT, conf)
	if err != nil {
		return nil, fmt.Errorf("oauth2 client: error exchanging Okta JWT for bearer token: %w", err)
	}

	tokenSource := &customTokenSource{
		conf:    conf,
		ctx:     ctx,
		oktaJWK: oktaJWK,
		token: &oauth2.Token{
			AccessToken: token.AccessToken,
		},
	}
	// reuse the tokenSource to refresh the token (automatically calls the custom Token() method when token is no longer valid).
	client := oauth2.NewClient(ctx, oauth2.ReuseTokenSource(token, tokenSource))

	return client, nil
}

// Token implements the oauth2.TokenSource interface and helps to implement custom token refresh logic.
// parent context is passed via the customTokenSource struct since we cannot modify the function signature here.
func (ts *customTokenSource) Token() (*oauth2.Token, error) {
	ts.mu.Lock()
	oktaJWT, err := generateOktaJWT(ts.oktaJWK, ts.conf)
	if err != nil {
		return nil, fmt.Errorf("error generating Okta JWT: %w", err)
	}
	token, err := exchangeForBearerToken(ts.ctx, oktaJWT, ts.conf)
	if err != nil {
		return nil, fmt.Errorf("error exchanging Okta JWT for bearer token: %w", err)

	}
	ts.mu.Unlock()
	return token, nil
}

func generateOktaJWT(oktaJWK string, cnf *oauth2.Config) (string, error) {
	// unmarshal the JWK into a map
	var jwkData map[string]string
	err := json.Unmarshal([]byte(oktaJWK), &jwkData)
	if err != nil {
		return "", fmt.Errorf("error decoding JWK: %w", err)
	}

	// create a RSA private key from JWK components
	decodeBase64 := func(key string) (*big.Int, error) {
		data, err := base64.RawURLEncoding.DecodeString(jwkData[key])
		if err != nil {
			return nil, fmt.Errorf("error decoding RSA JWK component %s: %w", key, err)
		}
		return new(big.Int).SetBytes(data), nil
	}

	n, err := decodeBase64("n")
	if err != nil {
		return "", err
	}
	e, err := decodeBase64("e")
	if err != nil {
		return "", err
	}
	d, err := decodeBase64("d")
	if err != nil {
		return "", err
	}
	p, err := decodeBase64("p")
	if err != nil {
		return "", err
	}
	q, err := decodeBase64("q")
	if err != nil {
		return "", err
	}
	dp, err := decodeBase64("dp")
	if err != nil {
		return "", err
	}
	dq, err := decodeBase64("dq")
	if err != nil {
		return "", err
	}
	qi, err := decodeBase64("qi")
	if err != nil {
		return "", err
	}

	privateKeyRSA := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: n,
			E: int(e.Int64()),
		},
		D:      d,
		Primes: []*big.Int{p, q},
		Precomputed: rsa.PrecomputedValues{
			Dp:   dp,
			Dq:   dq,
			Qinv: qi,
		},
	}

	// create a JWT token using required claims and sign it with the private key
	tok, err := jwt.NewBuilder().Audience([]string{cnf.Endpoint.TokenURL}).
		Issuer(cnf.ClientID).
		Subject(cnf.ClientID).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(time.Hour)).
		Build()
	if err != nil {
		return "", err
	}
	signedToken, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, privateKeyRSA))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return string(signedToken), nil
}

// exchangeForBearerToken exchanges the Okta JWT for a bearer token.
func exchangeForBearerToken(ctx context.Context, oktaJWT string, cnf *oauth2.Config) (*oauth2.Token, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", strings.Join(cnf.Scopes, " "))
	data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Set("client_assertion", oktaJWT)
	oauthConfig := &clientcredentials.Config{
		TokenURL:       cnf.Endpoint.TokenURL,
		EndpointParams: data,
	}
	tokenSource := oauthConfig.TokenSource(ctx)

	// get the access token
	accessToken, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
