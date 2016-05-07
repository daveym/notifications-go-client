package client

// IGovNotifyAPI - Client to handle communication with Gov Notify API.
type IGovNotifyAPI interface {
	Send(request GovNotifyHTTPClientRequest) (GovNotifyHTTPClientResponse, error)
}

// GovNotifyAPI - Implementation for IGovNotifyAPI interface
type GovNotifyAPI struct {
}

// Send -  Sends a request to GovNotify
func (p *GovNotifyAPI) Send(request GovNotifyHTTPClientRequest) (GovNotifyHTTPClientResponse, error) {

	var res GovNotifyHTTPClientResponse
	var err error

	return res, err

}
