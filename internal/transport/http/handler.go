package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comment service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

//SetupRoutes - set up all the roues for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Healthy")
	})
}
