package client

// IGovNotifyHTTPClientRequest - Model for HTTP request to GovNotify
type IGovNotifyHTTPClientRequest interface {
	GetMethod() HTTPMethod
	GetURL() string
	GetBody() string
	GetHeaders() map[string]string
	Build()
}

// GovNotifyHTTPClientRequest - Model for HTTP request to GovNotify
type GovNotifyHTTPClientRequest struct {
	_httpmethod HTTPMethod
	_url        string
	_body       string
	_headers    map[string]string
}
