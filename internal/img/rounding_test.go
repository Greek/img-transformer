package img

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/png"
	"testing"

	"github.com/greek/img-transform/internal/lib"
)

func TestApplyRounding(t *testing.T) {
	// Create a dummy image
	width := 100
	height := 100
	src := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			src.Set(x, y, color.RGBA{255, 0, 0, 255}) // Red
		}
	}

	// Encode the image to a buffer
	inputBuf := new(bytes.Buffer)
	err := png.Encode(inputBuf, src)
	if err != nil {
		t.Fatalf("failed to encode image: %v", err)
	}

	// Apply rounding
	outputBuf := new(bytes.Buffer)
	radius := "20"
	err = applyRounding(inputBuf, outputBuf, radius)
	if err != nil {
		t.Fatalf("applyRounding failed: %v", err)
	}

	// Decode the resulting image
	resultImg, _, err := image.Decode(outputBuf)
	if err != nil {
		t.Fatalf("failed to decode rounded image: %v", err)
	}

	// Check corner pixels - they should be transparent
	corners := []image.Point{
		{0, 0},
		{width - 1, 0},
		{0, height - 1},
		{width - 1, height - 1},
	}

	for _, p := range corners {
		c := resultImg.At(p.X, p.Y)
		_, _, _, a := c.RGBA()
		if a != 0 {
			t.Errorf("Expected corner pixel at %v to be transparent, but got alpha %d", p, a)
		}
	}

	// Check center pixel - it should be opaque
	center := image.Point{width / 2, height / 2}
	c := resultImg.At(center.X, center.Y)
	_, _, _, a := c.RGBA()
	if a == 0 {
		t.Errorf("Expected center pixel to be opaque, but it was transparent")
	}
}

func TestApplyRounding_Error(t *testing.T) {
	inputBuf := new(bytes.Buffer)
	outputBuf := new(bytes.Buffer)
	radius := "513" // > 512

	err := applyRounding(inputBuf, outputBuf, radius)

	var roundErr lib.ErrResponse
	if !errors.As(err, &roundErr) {
		t.Fatalf("Expected a RoundErr, but got %T", err)
	}

	if roundErr.ErrCode() != 400 {
		t.Errorf("Expected error code 400, but got %d", roundErr.ErrCode())
	}

	expectedReason := "radius too large"
	if roundErr.ErrReason() != expectedReason {
		t.Errorf("Expected reason '%s', but got '%s'", expectedReason, roundErr.ErrReason())
	}
}
