package model

import "regexp"
import "github.com/pkg/errors"

// IEmailRequest - Represents an email request interface
type IEmailRequest interface {
	GetEmail() string
	GetTemplateId() string
	GetPersonalisation() Personalisation
	Build(string, string, Personalisation) (EmailRequest, error)
}

// EmailRequest - implementation of IEmailRequest interface
type EmailRequest struct {
	_email           string
	_templateID      string
	_personalisation Personalisation
}

// GetEmail - Returns the email address.
func (p *EmailRequest) GetEmail() string {
	return p._email
}

// GetTemplateID - Returns a TemplateID
func (p *EmailRequest) GetTemplateID() string {
	return p._templateID
}

// GetPersonalisation - Returns a Personalisation struct
func (p *EmailRequest) GetPersonalisation() Personalisation {
	return p._personalisation
}

// Build - Builds a new Email Request object
func (p *EmailRequest) Build(email string, templateID string, pers Personalisation) (EmailRequest, error) {

	var req EmailRequest
	var err error

	if validateEmail(email) {
		errors.Wrap(err, "Email validation failed")
		return req, err
	}

	if templateID == "" {
		errors.Wrap(err, "Template id cannot be empty")
		return req, err
	}

	req._email = email
	req._templateID = templateID
	req._personalisation = pers

	return req, err
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
