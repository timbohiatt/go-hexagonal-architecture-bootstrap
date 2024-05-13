package http

import (
	"fmt"
	"net/http"
)

// HTTP REST Endpoint Handlers
func (a Adapter) HandleHealthzGet(w http.ResponseWriter, r *http.Request) error {
	// Declare Response Struct
	type HandleResponse struct {
		StatusCode int    `json:"statusCode"`
		StatusMsg  string `json:"statusMsg"`
	}

	// Execute Health Check
	fmt.Println("Made it to HTTP healthz")
	_, err := a.api.Healthz()
	if err != nil {
		// Error - Generic - return http 500
		WriteJSON(w, http.StatusOK, &HandleResponse{
			StatusCode: 500,
			StatusMsg:  "service is encountering issues",
		})
	}

	// Success - return http 200 ok
	WriteJSON(w, http.StatusOK, &HandleResponse{
		StatusCode: 200,
		StatusMsg:  "service is healthy",
	})

	return nil
}

// Handler Supporting Functions
type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, apiError{
				Error: err.Error(),
			})
		}
	}
}

// Handler Supporting Types

type apiError struct {
	Error string
}
