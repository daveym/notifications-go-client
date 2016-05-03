package model

import "github.com/pkg/errors"

// IPersonalisation - Interface representing information used to personalize a message sent through GovNotify service.
type IPersonalisation interface {
	asMap() string
	Build(string, string) (Personalisation, error)
}

// Personalisation - Implementation of the IPersonalisation interface
type Personalisation struct {
	_store map[string]string
}

// asMap - returns the personalisation information as a map
func (p *Personalisation) asMap() map[string]string {
	return p._store
}

// Build - Builds a new Personalisation instance
func (p *Personalisation) Build(name string, value string) (Personalisation, error) {

	var res Personalisation
	var err error

	if name == "" || value == "" {
		errors.Wrap(err, "Personalisation when set must have at least one field")
		return res, err
	}

	p._store = make(map[string]string)
	p._store[name] = value

	return res, err
}
