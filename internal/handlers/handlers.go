package handlers

import (
	"github.com/gorilla/mux"
	"github.com/greek/img-transform/internal/handlers/files"
	"github.com/greek/img-transform/internal/handlers/root"
)

// Register initializes all the handlers for the server.
func Register(r *mux.Router) {
	root.RegisterRoutes(r)
	files.RegisterRoutes(r)
}
