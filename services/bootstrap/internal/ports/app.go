package ports

type APIPort interface {
	Healthz() (bool, error)
}
