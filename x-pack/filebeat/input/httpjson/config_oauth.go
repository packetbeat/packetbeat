// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package httpjson

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/endpoints"
	"golang.org/x/oauth2/google"
)

type OAuth2Provider string

const (
	OAuth2ProviderDefault OAuth2Provider = ""
	OAuth2ProviderAzure   OAuth2Provider = "azure"
	OAuth2ProviderGoogle  OAuth2Provider = "google"
)

func (p OAuth2Provider) canonical() OAuth2Provider {
	return OAuth2Provider(strings.ToLower(string(p)))
}

// OAuth2 contains information about oauth2 authentication settings.
type OAuth2 struct {
	// common oauth fields
	ClientID       string              `config:"client.id"`
	ClientSecret   string              `config:"client.secret"`
	Enabled        *bool               `config:"enabled"`
	EndpointParams map[string][]string `config:"endpoint_params"`
	Provider       string              `config:"provider"`
	Scopes         []string            `config:"scopes"`
	TokenURL       string              `config:"token_url"`

	// google specific
	GoogleCredentialsFile string `config:"google.credentials_file"`
	GoogleCredentialsJSON []byte `config:"google.credentials_json"`
	GoogleJWTFile         string `config:"google.jwt_file"`

	// microsoft azure specific
	AzureTenantID string `config:"azure.tenant_id"`
}

// IsEnabled returns true if the `enable` field is set to true in the yaml.
func (o *OAuth2) IsEnabled() bool {
	return o != nil && (o.Enabled == nil || *o.Enabled)
}

func (o *OAuth2) Client(ctx context.Context, client *http.Client) (*http.Client, error) {
	ctx = context.WithValue(ctx, oauth2.HTTPClient, client)

	switch o.GetProvider() {
	case OAuth2ProviderGoogle:
		creds, err := google.CredentialsFromJSON(ctx, o.GoogleCredentialsJSON, o.Scopes...)
		if err != nil {
			return nil, fmt.Errorf("oauth2 client: error loading credentials: %w", err)
		}
		return oauth2.NewClient(ctx, creds.TokenSource), nil
	}

	creds := clientcredentials.Config{
		ClientID:       o.ClientID,
		ClientSecret:   o.ClientSecret,
		TokenURL:       o.GetTokenURL(),
		Scopes:         o.Scopes,
		EndpointParams: o.EndpointParams,
	}

	return creds.Client(ctx), nil
}

func (o *OAuth2) GetTokenURL() string {
	switch o.GetProvider() {
	case OAuth2ProviderAzure:
		if o.TokenURL == "" {
			return endpoints.AzureAD(o.AzureTenantID).TokenURL
		}
	}

	return o.TokenURL
}

func (o OAuth2) GetProvider() OAuth2Provider {
	return OAuth2Provider(o.Provider).canonical()
}

func (o *OAuth2) Validate() error {
	switch o.GetProvider() {
	case OAuth2ProviderAzure:
		return o.validateAzureProvider()
	case OAuth2ProviderGoogle:
		return o.validateGoogleProvider()
	case OAuth2ProviderDefault:
		if o.TokenURL == "" || o.ClientID == "" || o.ClientSecret == "" {
			return errors.New("invalid configuration: both token_url and client credentials must be provided")
		}
	default:
		return fmt.Errorf("invalid configuration: unknown provider %q", o.GetProvider())
	}
	return nil
}

func (o *OAuth2) validateGoogleProvider() error {
	if o.TokenURL != "" || o.ClientID != "" || o.ClientSecret != "" ||
		o.AzureTenantID != "" || len(o.EndpointParams) > 0 {
		return errors.New("invalid configuration: none of token_url and client credentials can be used, use google.credentials_file, google.jwt_file, google.credentials_json or ADC instead")
	}

	// credentials_json
	if len(o.GoogleCredentialsJSON) > 0 {
		return nil
	}

	// credentials_file
	if o.GoogleCredentialsFile != "" {
		return o.populateCredentialsJSONFromFile(o.GoogleCredentialsFile)
	}

	// jwt_file
	if o.GoogleJWTFile != "" {
		return o.populateCredentialsJSONFromFile(o.GoogleJWTFile)
	}

	// Application Default Credentials (ADC)
	ctx := context.Background()
	if creds, err := google.FindDefaultCredentials(ctx, o.Scopes...); err == nil {
		o.GoogleCredentialsJSON = creds.JSON
		return nil
	}

	return fmt.Errorf("invalid configuration: no authentication credentials were configured or detected (ADC)")
}

func (o *OAuth2) populateCredentialsJSONFromFile(file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return fmt.Errorf("invalid configuration: the file %q cannot be found", file)
	}

	credBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("invalid configuration: the file %q cannot be read", file)
	}

	o.GoogleCredentialsJSON = credBytes

	return nil
}

func (o *OAuth2) validateAzureProvider() error {
	if o.TokenURL == "" && o.AzureTenantID == "" {
		return errors.New("invalid configuration: at least one of token_url or tenant_id must be provided")
	}
	if o.TokenURL != "" && o.AzureTenantID != "" {
		return errors.New("invalid configuration: only one of token_url and tenant_id can be used")
	}
	if o.ClientID == "" || o.ClientSecret == "" {
		return errors.New("invalid configuration: client credentials must be provided")
	}

	return nil
}
