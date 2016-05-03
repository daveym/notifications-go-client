package model

// INotificationCreatedResponse - Interface that represents information available at the time of requesting a notification.
type INotificationCreatedResponse interface {
	GetID() string
	Build(string) (NotificationCreatedResponse, error)
}

// NotificationCreatedResponse - Implementation of INotificationCreatedResponse interface
type NotificationCreatedResponse struct {
	_ID string
}

// GetID - Returns the notification response ID.
func (p *NotificationCreatedResponse) GetID() string {
	return p._ID
}

// Build - Builds a new instance of a NotificationCreatedResponse
func (p *NotificationCreatedResponse) Build(ID string) (NotificationCreatedResponse, error) {

	var res NotificationCreatedResponse
	var err error

	res._ID = ID

	return res, err
}
