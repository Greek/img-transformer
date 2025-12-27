package files

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/{bucket}/{key}", GetFile).Methods("GET")
}

// TransformationCmds
type TransformationCmds struct {
}

func parseTransformation(fragment string) []string {
	_, aft, _ := strings.Cut(fragment, "=")
	commands := strings.Split(aft, ",")

	return commands
}

// GetFile retrieves a file from a specified bucket.
func GetFile(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	bucket := vars["bucket"]
	key, _, _ := strings.Cut(vars["key"], "=")
	transforms := parseTransformation(req.URL.Path)

	fmt.Println(vars)
	fmt.Println(bucket)
	fmt.Println(key)
	fmt.Println(transforms)

	w.Write([]byte("Hello world! again..."))
}
