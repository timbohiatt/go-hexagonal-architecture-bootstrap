package ports

import "sync"

type HTTPPort interface {
	Run(hostname string, port string, wg *sync.WaitGroup)
}
