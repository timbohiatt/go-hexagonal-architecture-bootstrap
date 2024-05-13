package core

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) Healthz() (bool, error) {
	return true, nil
}
