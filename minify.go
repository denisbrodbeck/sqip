package sqip

import (
	"bytes"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/svg"
)

// Minify takes a svg and returns a minified version of the input.
func Minify(in string) (out string, err error) {
	reader := strings.NewReader(in)
	writer := bytes.NewBuffer(nil)

	min := minify.New()
	min.AddFunc("image/svg+xml", svg.Minify)

	if err := min.Minify("image/svg+xml", writer, reader); err != nil {
		return "", err
	}
	return writer.String(), nil
}
