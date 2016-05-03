package model

import "github.com/pkg/errors"

// IStatusRequest - Interface defining the  id of notification status
type IStatusRequest interface {
	GetNotificationID() string
}

// StatusRequest - Implementation of SItatusRequest interface
type StatusRequest struct {
	_notificationID string
}

// GetNotificationID - Returns a notification ID
func (p *StatusRequest) GetNotificationID() string {
	return p._notificationID
}

// Build - Builds a new instance of a StatusReuest
func (p *StatusRequest) Build(notificationID string) (StatusRequest, error) {

	var req StatusRequest
	var err error

	if notificationID == "" {
		errors.Wrap(err, "notificationId cannot be empty")
		return req, err
	}

	req._notificationID = notificationID

	return req, err
}
