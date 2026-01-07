package img

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"strings"

	"github.com/greek/img-transform/internal/lib"
	"github.com/greek/img-transform/internal/lib/logging"
)

// ApplyTransformations transforms a given image using provided valid
// image transformation commands.
func ApplyTransformations(imageReader io.Reader, transformations []string) (io.Reader, error) {
	logger := logging.BuildLogger("ApplyTransformations")

	var err error
	currentImage := imageReader

	for _, v := range transformations {
		logger.Info("Applying transformation: " + v)

		processedImage := new(bytes.Buffer)
		writer := io.Writer(processedImage)

		switch bef, _, _ := strings.Cut(v, "_"); bef {
		case "round":
			{
				radiusStr := strings.Split(v, "_")[1]
				err = applyRounding(currentImage, writer, radiusStr)
			}
		}

		// needed to guarantee image is processed regardless of any transformations
		logger.Debug("copying current image onto writer")
		_, err = io.Copy(writer, currentImage)

		if err != nil {
			var ogErr lib.ErrResponse
			if errors.As(err, &ogErr) {
				logger.Error("Failed to transform image", slog.String("cause", ogErr.ErrReason()))
				return nil, err
			}
		}

		currentImage = bytes.NewReader(processedImage.Bytes())
	}

	return currentImage, nil
}
