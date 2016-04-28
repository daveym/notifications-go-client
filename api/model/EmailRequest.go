package model

import "regexp"

// EmailRequest - Represents an email request
type EmailRequest interface {
	GetEmail() string
	GetTemplateId() string
	GetPersonalisation() personalisation
}

type emailRequest struct {
	_email           string
	_templateID      string
	_personalisation personalisation
}

func (p *emailRequest) GetEmail() string {
	return p._email
}

func (p *emailRequest) GetTemplateID() string {
	return p._templateID
}

func (p *emailRequest) GetPersonalisation() personalisation {
	return p._personalisation
}

func (p *emailRequest) Build(email string, templateID string, pers personalisation) (emailRequest, error) {

	var req emailRequest
	var err error

	if validateEmail(email) {
		req._email = email
		req._templateID = templateID
		req._personalisation = pers

	}

	return req, err
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
