package model

import "github.com/pkg/errors"

// Personalisation - Represents information used to personalize a message sent through GovNotify service.
type Personalisation interface {
	asMap() string
	Build(string, string) (personalisation, error)
}

type personalisation struct {
	_store map[string]string
}

func (p *personalisation) asMap() map[string]string {
	return p._store
}

func (p *personalisation) Build(name string, value string) (personalisation, error) {

	var res personalisation
	var err error

	if name == "" || value == "" {
		errors.Wrap(err, "Personalisation when set must have at least one field")
		return res, err
	}

	p._store = make(map[string]string)
	p._store[name] = value

	return res, err
}
