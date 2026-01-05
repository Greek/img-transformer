package img

import (
	"bytes"
	"io"
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

		// Create a buffer to hold the output of the transformation
		processedImage := new(bytes.Buffer)
		writer := io.Writer(processedImage)

		switch bef, _, _ := strings.Cut(v, "_"); bef {
		case "round":
			{
				radiusStr := strings.Split(v, "_")[1]
				err = applyRounding(currentImage, writer, radiusStr)
			}
		default:
			return nil, &lib.HTTPErr{Reason: "transformation not supported", Code: 400}
		}

		if err != nil {
			return nil, err
		}

		// The output of the last transformation becomes the input for the next
		currentImage = bytes.NewReader(processedImage.Bytes())
	}

	return currentImage, nil
}
