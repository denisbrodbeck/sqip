package sqip

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"  // register gif
	_ "image/jpeg" // register jpeg
	_ "image/png"  // register png
	"io"
	"io/ioutil"
	"os"
)

// LoadImage takes a path and returns the content as an image.
func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return decodeImage(file)
}

func decodeImage(r io.Reader) (image.Image, error) {
	im, _, err := image.Decode(r)
	return im, err
}

// SaveFile creates the named file with mode 0666 (before umask), truncating
// it if it already exists, and writes the content to it.
func SaveFile(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0666)
}

// StringBase64 is a base64 encoded string.
type StringBase64 string

// Base64 returns the base64 encoding of src.
func Base64(src string) StringBase64 {
	enc := base64.StdEncoding.EncodeToString([]byte(src))
	return StringBase64(enc)
}

// ImageTag creates an example <img> tag.
func ImageTag(filename string, svg StringBase64, width, height int) string {
	return fmt.Sprintf(`<img width="%d" height="%d" src="%s" alt="Add descriptive alt text" style="background-size: cover; background-image: url(data:image/svg+xml;base64,%s);">`, width, height, filename, svg)
}

// ImageWidthAndHeight returns the width and height of input image.
func ImageWidthAndHeight(input image.Image) (width, height int) {
	bounds := input.Bounds()
	return bounds.Dx(), bounds.Dy()
}

// largerOne returns the larger int of two ints.
func largerOne(x, y int) int {
	if x > y {
		return x
	}
	return y
}
