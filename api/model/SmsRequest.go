package model

// SmsRequest - Represents an SMS request
type SmsRequest interface {
	getPhoneNumber() string
	getTemplateId() string
	getPersonalisation() personalisation
}

type smsRequest struct {
	_phonenumber     string
	_templateID      string
	_personalisation personalisation
}

func (p *smsRequest) getPhoneNumber() string {
	return p._phonenumber
}

func (p *smsRequest) getTemplateID() string {
	return p._templateID
}

func (p *smsRequest) getPersonalisation() personalisation {
	return p._personalisation
}
