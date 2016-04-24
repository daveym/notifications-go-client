package model

// EmailRequest - Represents an email request
type EmailRequest interface {
	getEmail() string
	getTemplateId() string
	getPersonalisation() personalisation
}

type emailRequest struct {
	_email           string
	_templateID      string
	_personalisation personalisation
}

type personalisation struct {
}

func (p *emailRequest) getEmail() string {
	return p._email
}

func (p *emailRequest) getTemplateID() string {
	return p._templateID
}

func (p *emailRequest) getPersonalisation() personalisation {
	return p._personalisation
}
