package files

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	s3lib "github.com/greek/img-transform/internal/lib/s3"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/{bucket}/{key}", GetFile).Methods("GET")
}

var s3 = s3lib.InitS3()

// TransformationCmds
type MockData struct {
	Transformations []string `json:"commands"`
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

	mockData := MockData{Transformations: transforms}

	_, err := s3.GetFile(bucket, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(mockData); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
