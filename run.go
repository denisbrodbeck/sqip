package sqip

// Run takes a file and primitve related config properties and creates a SVG-based LQIP image.
func Run(file string, workSize, count, mode, alpha, repeat, workers int, background string) (out string, width, height int, err error) {
	// Load image
	image, err := LoadImage(file)
	if err != nil {
		return "", 0, 0, err
	}
	// Use image-size to retrieve the width and height dimensions of the input image
	// We need these sizes to pass to Primitive and to write the SVG viewbox
	w, h := ImageWidthAndHeight(image)
	// Since Primitive is only interested in the larger dimension of the input image, let's find it
	outputSize := largerOne(w, h)

	// create primitive
	svg, err := Primitive(image, workSize, outputSize, count, mode, alpha, repeat, workers, background)
	if err != nil {
		return "", 0, 0, err
	}

	// minify svg
	svg, err = Minify(svg)
	if err != nil {
		return "", 0, 0, err
	}

	// blur the svg
	svg, err = Blur(svg, w, h)
	if err != nil {
		return "", 0, 0, err
	}
	return svg, w, h, nil
}
