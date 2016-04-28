package model

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

type personalisation struct {
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

func (p *emailRequest) validate() bool {
	// TODO VALIDATE EMAIL

	return true
}

func (p *emailRequest) Build(email string) emailRequest {

	var req emailRequest
	req._email = email

	return req

}
