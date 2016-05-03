package model

// IStatusResponse - Interface Representing status information of the already requested notifications.
type IStatusResponse interface {
	GetID() string
	GetStatus() string
	Build(string, string) (StatusResponse, error)
}

// StatusResponse - Implementation of IStatusResponse Interface
type StatusResponse struct {
	_ID     string
	_status string
}

// GetID - Returns a status response ID
func (p *StatusResponse) GetID() string {
	return p._ID
}

// GetStatus - Returns a status response status
func (p *StatusResponse) GetStatus() string {
	return p._status
}

// Build - Builds a new instance of a StatusResponse
func (p *StatusResponse) Build(id string, status string) (StatusResponse, error) {

	var res StatusResponse
	var err error

	res._ID = id
	res._status = status

	return res, err
}
