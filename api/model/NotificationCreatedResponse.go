package model

// NotificationCreatedResponse - Represents information available at the time of requesting a notification.
type NotificationCreatedResponse interface {
	GetID() string
}

type notificationCreatedResponse struct {
	_ID string
}

func (p *notificationCreatedResponse) GetID() string {
	return p._ID
}

func (p *notificationCreatedResponse) Build(ID string) (notificationCreatedResponse, error) {

	var res notificationCreatedResponse
	var err error

	res._ID = ID

	return res, err
}
