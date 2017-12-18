package sqip

import (
	"testing"
)

func TestPrimitive(t *testing.T) {
	img, err := LoadImage(testImage)
	if err != nil {
		t.Error(err)
	}
	w, h := ImageWidthAndHeight(img)
	outputSize := largerOne(w, h)
	svg, err := Primitive(img, 64, outputSize, 8, 0, 128, 0, 1, "")
	if err != nil {
		t.Error(err)
	}
	if svg == "" {
		t.Error("Primitive() output is empty")
	}
}
