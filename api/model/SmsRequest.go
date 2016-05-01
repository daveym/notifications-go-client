package model

import "github.com/pkg/errors"

// SmsRequest - Represents an SMS request
type SmsRequest interface {
	GetEmail() string
	GetTemplateId() string
	GetPersonalisation() personalisation
	Build(string, string, personalisation) (smsRequest, error)
}

type smsRequest struct {
	_phoneNumber     string
	_templateID      string
	_personalisation personalisation
}

func (p *smsRequest) GetPhoneNumber() string {
	return p._phoneNumber
}

func (p *smsRequest) GetTemplateID() string {
	return p._templateID
}

func (p *smsRequest) GetPersonalisation() personalisation {
	return p._personalisation
}

func (p *smsRequest) Build(phoneNumber string, templateID string, pers personalisation) (smsRequest, error) {

	var req smsRequest
	var err error

	if phoneNumber == "" {
		errors.Wrap(err, "Phone number cannot be empty")
		return req, err
	}

	if templateID == "" {
		errors.Wrap(err, "Template id cannot be empty")
		return req, err
	}

	req._phoneNumber = phoneNumber
	req._templateID = templateID
	req._personalisation = pers

	return req, err
}
