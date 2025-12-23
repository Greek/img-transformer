package files

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/{bucket}/{key}", GetFile).Methods("GET")
}

// GetFile retrieves a file from a specified bucket.
func GetFile(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Println(vars)

	w.Write([]byte("Hello world! again..."))
}
