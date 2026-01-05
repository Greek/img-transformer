package img

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"strconv"

	"github.com/greek/img-transform/internal/lib"
)

func applyRounding(r io.Reader, w io.Writer, radiusStr string) error {
	radius, err := strconv.Atoi(radiusStr)
	if radius > 512 {
		return &lib.HTTPErr{Code: 400, Reason: "radius too large"}
	}

	img, _, err := image.Decode(r)
	if err != nil {
		return &lib.HTTPErr{Code: 500, Reason: "image decoding failed"}
	}

	// Create a new RGBA image with the same bounds as the original
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	// Create a rounded rectangle mask
	mask := &rounded{
		p: image.Point{X: bounds.Dx(), Y: bounds.Dy()},
		r: radius,
	}

	// Draw the original image onto the destination image using the mask
	draw.DrawMask(dst, bounds, img, image.Point{}, mask, image.Point{}, draw.Src)

	// Encode the result to the writer as a PNG
	err = png.Encode(w, dst)
	if err != nil {
		return &lib.HTTPErr{Code: 500, Reason: "image encoding failed"}
	}
	return nil
}

// rounded is a shape that implements the image.Image interface
// and represents a rounded rectangle.
type rounded struct {
	p image.Point
	r int
}

func (c *rounded) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *rounded) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.p.X, c.p.Y)
}

func (c *rounded) At(x, y int) color.Color {
	// Check if the point is within the rounded corners
	// top-left
	if x < c.r && y < c.r && (x-c.r)*(x-c.r)+(y-c.r)*(y-c.r) > c.r*c.r {
		return color.Alpha{0}
	}
	// top-right
	if x > c.p.X-c.r && y < c.r && (x-(c.p.X-c.r))*(x-(c.p.X-c.r))+(y-c.r)*(y-c.r) > c.r*c.r {
		return color.Alpha{0}
	}
	// bottom-left
	if x < c.r && y > c.p.Y-c.r && (x-c.r)*(x-c.r)+(y-(c.p.Y-c.r))*(y-(c.p.Y-c.r)) > c.r*c.r {
		return color.Alpha{0}
	}
	// bottom-right
	if x > c.p.X-c.r && y > c.p.Y-c.r && (x-(c.p.X-c.r))*(x-(c.p.X-c.r))+(y-(c.p.Y-c.r))*(y-(c.p.Y-c.r)) > c.r*c.r {
		return color.Alpha{0}
	}
	return color.Alpha{255}
}
