package model

// StatusResponse - Represents status information of the already requested notifications.
type StatusResponse interface {
	GetID() string
	GetStatus() string
	Build(string, string) (statusResponse, error)
}

type statusResponse struct {
	_ID     string
	_status string
}

func (p *statusResponse) GetID() string {
	return p._ID
}

func (p *statusResponse) GetStatus() string {
	return p._status
}

func (p *statusResponse) Build(id string, status string) (statusResponse, error) {

	var res statusResponse
	var err error

	res._ID = id
	res._status = status

	return res, err
}
