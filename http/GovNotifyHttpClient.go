package client

// GovNotifyAPI - Client to handle communication with Gov Notify API.
type GovNotifyAPI interface {
	send(request GovNotifyHTTPClientRequest) (GovNotifyHTTPClientResponse, error)
}
