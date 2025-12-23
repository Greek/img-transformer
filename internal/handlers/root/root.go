package root

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", GetRoot).Methods("GET")
}

func GetRoot(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world!"))
}
