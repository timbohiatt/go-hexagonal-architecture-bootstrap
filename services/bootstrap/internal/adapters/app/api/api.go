package api

import (
	"persona/internal/ports"
)

type Adapter struct {
	core ports.CorePort
}

func NewAdapter(core ports.CorePort) *Adapter {
	return &Adapter{
		core: core,
	}
}

func (a Adapter) Healthz() (bool, error) {
	return a.core.Healthz()
}
