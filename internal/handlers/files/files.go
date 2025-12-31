package files

import (
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/greek/img-transform/internal/img"
	"github.com/greek/img-transform/internal/lib/logging"
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
	logger := logging.BuildLogger("GetFile")
	vars := mux.Vars(req)

	bucket := vars["bucket"]
	key, _, _ := strings.Cut(vars["key"], "=")
	transforms := parseTransformation(req.URL.Path)

	logger.Info("Getting file "+bucket+"/"+key, slog.Any("transforms", transforms))

	data, err := s3.GetFile(bucket, key)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer data.Body.Close()

	img.ApplyTransformations(&data.Body, transforms)

	w.Header().Set("Content-Type", *data.ContentType)
	if _, err := io.Copy(w, data.Body); err != nil {
		w.WriteHeader(404)
		w.Write([]byte("unable to write file data"))
	}
}
