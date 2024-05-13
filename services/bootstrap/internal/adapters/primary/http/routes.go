package http

import (
	"github.com/go-chi/chi/v5"
)

func addRoutes(a Adapter, router *chi.Mux) {
	router.Get("/api/v1/healthz", makeHandlerFunc(a.HandleHealthzGet))
}
