package geom

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/go-gl/mathgl/mgl32"
)

type Image interface {
	Geometry2D

	// RGBA converts the image to an RGBA image, return nil if image cannot be converted
	RGBA() *image.RGBA
}

type imageStruct struct {
	geometry

	data  image.Image
	valid bool
}

// NewImage creates a new image based on the reader. Supported formats are JPG and PNG
func NewImage(r io.Reader) Image {

	var err error
	img := imageStruct{
		valid: false,
		geometry: geometry{
			dimension: 2,
		},
	}

	img.data, _, err = image.Decode(r)
	if err != nil {
		return &img
	}
	img.valid = true
	return &img
}

func (i *imageStruct) Valid() bool {
	return i.valid
}

func (i *imageStruct) Width() int {
	return i.data.Bounds().Size().X
}

func (i *imageStruct) Height() int {
	return i.data.Bounds().Size().Y
}

func (i *imageStruct) RGBA() *image.RGBA {
	rgba := image.NewRGBA(i.data.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil
	}
	draw.Draw(rgba, rgba.Bounds(), i.data, image.Point{0, 0}, draw.Src)
	return rgba
}

func (i *imageStruct) MinBound() mgl32.Vec2 {
	return mgl32.Vec2{
		float32(i.data.Bounds().Min.X),
		float32(i.data.Bounds().Min.Y),
	}
}

func (i *imageStruct) MaxBound() mgl32.Vec2 {
	return mgl32.Vec2{
		float32(i.data.Bounds().Max.X),
		float32(i.data.Bounds().Max.Y),
	}
}
