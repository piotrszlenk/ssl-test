package endpoint

import (
	"github.com/piotrszlenk/ssl-test/pkg/logz"
)

type Endpoints struct {
	Items  []Endpoint
	logger logz.LogHandler
	path   string
}

type Endpoint struct {
	Fqdn string
	Port uint16
}

func NewEndpoints(path string, logger logz.LogHandler) *Endpoints {
	e := new(Endpoints)
	e.logger = logger
	e.path = path
	return e
}

func (e *Endpoints) LoadEndpoints() (Endpoints, error) {
	return *e, nil
}
