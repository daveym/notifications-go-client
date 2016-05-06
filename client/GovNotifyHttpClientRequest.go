package client

import "github.com/pkg/errors"

// IGovNotifyHTTPClientRequest - Model for HTTP request to GovNotify
type IGovNotifyHTTPClientRequest interface {
	GetMethod() HTTPMethod
	GetURL() string
	GetBody() string
	GetHeaders() map[string]string
	Build(HTTPMethod, string, string, map[string]string) (GovNotifyHTTPClientRequest, error)
}

// GovNotifyHTTPClientRequest - Model for HTTP request to GovNotify
type GovNotifyHTTPClientRequest struct {
	_httpmethod HTTPMethod
	_url        string
	_body       string
	_headers    map[string]string
}

// GetMethod - Returns the HTTP Method
func (p *GovNotifyHTTPClientRequest) GetMethod() HTTPMethod {
	return p._httpmethod
}

// GetURL - Returns the Target URL
func (p *GovNotifyHTTPClientRequest) GetURL() string {
	return p._url
}

// GetBody - Returns the Request Body
func (p *GovNotifyHTTPClientRequest) GetBody() string {
	return p._body
}

// GetHeaders - Returns the Request Body
func (p *GovNotifyHTTPClientRequest) GetHeaders() map[string]string {
	return p._headers
}

// Build - Returns a new GovNotifyHTTPClientRequest
func (p *GovNotifyHTTPClientRequest) Build(method HTTPMethod, url string, body string, headers map[string]string) (GovNotifyHTTPClientRequest, error) {

	var req GovNotifyHTTPClientRequest
	var err error

	if method.value == "" {
		errors.Wrap(err, "Method must be specified")
		return req, err
	}

	if url == "" {
		errors.Wrap(err, "URL must be specified")
		return req, err
	}

	req._httpmethod = method
	req._url = url
	req._body = body
	req._headers = headers

	return req, err

}
