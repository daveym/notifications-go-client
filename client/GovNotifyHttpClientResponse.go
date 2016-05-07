package client

// IGovNotifyHTTPClientResponse - Model for HTTP response from GovNotify
type IGovNotifyHTTPClientResponse interface {
	GetBody() string
	GetStatusCode() int
	Build(string, int) (GovNotifyHTTPClientResponse, error)
}

// GovNotifyHTTPClientResponse - Model for HTTP response from GovNotify
type GovNotifyHTTPClientResponse struct {
	_body       string
	_statuscode int
}

// GetBody - Returns the response body
func (p *GovNotifyHTTPClientResponse) GetBody() string {
	return p._body
}

// GetStatusCode - Returns the response statuscode
func (p *GovNotifyHTTPClientResponse) GetStatusCode() int {
	return p._statuscode
}

// Build - Returns a new GovNotifyHTTPClientResponse
func (p *GovNotifyHTTPClientResponse) Build(body string, statuscode int) (GovNotifyHTTPClientResponse, error) {

	var res GovNotifyHTTPClientResponse
	var err error

	res._body = body
	res._statuscode = statuscode

	return res, err

}
