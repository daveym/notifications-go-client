package model

import "github.com/pkg/errors"

// ISMSRequest - Represents an SMS request interface
type ISMSRequest interface {
	GetEmail() string
	GetTemplateId() string
	GetPersonalisation() Personalisation
	Build(string, string, Personalisation) (SMSRequest, error)
}

// SMSRequest - Represents the implementation of the SMS request interface
type SMSRequest struct {
	_phoneNumber     string
	_templateID      string
	_personalisation Personalisation
}

// GetPhoneNumber - Return a phone number
func (p *SMSRequest) GetPhoneNumber() string {
	return p._phoneNumber
}

// GetTemplateID - Return the TemplateID
func (p *SMSRequest) GetTemplateID() string {
	return p._templateID
}

// GetPersonalisation - Return the personalisation
func (p *SMSRequest) GetPersonalisation() Personalisation {
	return p._personalisation
}

// Build - Build a new instance of an SMS Request
func (p *SMSRequest) Build(phoneNumber string, templateID string, pers Personalisation) (SMSRequest, error) {

	var req SMSRequest
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
