package model

import "github.com/pkg/errors"

// StatusRequest - Returns an id of notification a status of which is requested.
type StatusRequest interface {
	GetNotificationID() string
}

type statusRequest struct {
	_notificationID string
}

func (p *statusRequest) GetNotificationID() string {
	return p._notificationID
}

func (p *statusRequest) Build(notificationID string) (statusRequest, error) {

	var req statusRequest
	var err error

	if notificationID == "" {
		errors.Wrap(err, "notificationId cannot be empty")
		return req, err
	}

	req._notificationID = notificationID

	return req, err
}
