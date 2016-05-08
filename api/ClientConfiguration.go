package api

import (
	"net/url"

	"github.com/pkg/errors"
)

const keylength = 32

// IClientConfiguration - Interface to define configuration for API client.
type IClientConfiguration interface {
	getServiceID() string
	getSecret() string
	getBaseURL() string
	Build(string, string, string) (ClientConfiguration, error)
}

// ClientConfiguration - Implementation of IClientConfiguration
type ClientConfiguration struct {
	_serviceID string
	_secret    string
	_baseURL   string
}

// getServiceID - Returns the ServiceID
func (p *ClientConfiguration) getServiceID() string {
	return p._serviceID
}

// getSecret - Returns the secret
func (p *ClientConfiguration) getSecret() string {
	return p._secret
}

// getBaseURL - Returns the BaseURL
func (p *ClientConfiguration) getBaseURL() string {
	return p._baseURL
}

// Build - Returns a new ClientConfiguration
func (p *ClientConfiguration) Build(serviceID string, secret string, baseURL string) (ClientConfiguration, error) {

	var config ClientConfiguration
	var err error

	_, err = url.Parse(baseURL)
	if err != nil {
		errors.Wrap(err, "baseUrl should be a valid URL")
		return config, err
	}

	if len(secret) < keylength {
		errors.Wrap(err, "Secret should have at least characters: "+string(keylength))
		return config, err
	}

	if baseURL == "" {
		errors.Wrap(err, "ServiceId cannot be empty")
		return config, err
	}

	config._baseURL = baseURL
	config._secret = secret
	config._serviceID = serviceID

	return config, err

}
