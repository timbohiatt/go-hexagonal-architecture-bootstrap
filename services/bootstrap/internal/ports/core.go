package ports

type CorePort interface {
	Healthz() (bool, error)
}