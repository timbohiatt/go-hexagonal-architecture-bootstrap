package http

import (
	"persona/internal/ports"

	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (a Adapter) Run(hostname, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Check if the port is available for binding.
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("unable to listen on port %s: %v", port, err)
	}
	_ = ln.Close()

	router := chi.NewRouter()

	// Middleware: OOTB CHI Middleware;
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// Routes: Configure CHI HTTP Routes
	addRoutes(a, router)

	log.Printf("listening for REST API calls on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", hostname, port), router); err != nil {
		log.Fatalf("unable to listen on port %s: %v", port, err)
	}

	return

}
