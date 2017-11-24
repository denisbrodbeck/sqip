package sqip

import (
	"image"

	"github.com/fogleman/primitive/primitive"
	"github.com/nfnt/resize"
)

// Primitive generates a SVG consisting of several simple shapes
// that approximate the main features visible inside the image.
//
// All primitive generation related flags are supported.
// See https://github.com/fogleman/primitive for a list of supported flags
// and their meaning in respect to the output.
//
// You should seed your rng in your main function, before calling this function.
//
// Try:
//   rand.Seed(time.Now().UTC().UnixNano())
func Primitive(input image.Image, workSize, outputSize, count, mode, alpha, repeat, workers int, background string) (svg string, err error) {
	// scale down input image if needed
	if workSize > 0 {
		input = resize.Thumbnail(uint(workSize), uint(workSize), input, resize.Bilinear)
	}

	// determine background color
	var bg primitive.Color
	if background == "" {
		bg = primitive.MakeColor(primitive.AverageImageColor(input))
	} else {
		bg = primitive.MakeHexColor(background)
	}

	// run algorithm
	model := primitive.NewModel(input, bg, outputSize, workers)
	for i := 1; i <= count; i++ {
		// find optimal shape and add it to the model
		model.Step(primitive.ShapeType(mode), alpha, repeat)
	}
	return model.SVG(), nil
}
