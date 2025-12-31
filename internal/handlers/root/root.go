package root

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/greek/img-transform/internal/lib"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", GetRoot).Methods("GET")
}

func GetRoot(res http.ResponseWriter, req *http.Request) {
	data := map[string]string{"message": "Hello world!"}
	lib.WriteJSONSuccess(res, http.StatusOK, data)
}
