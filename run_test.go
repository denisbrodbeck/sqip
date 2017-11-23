package sqip

import "testing"

func TestRun(t *testing.T) {
	svg, w, h, err := Run(testImage, 64, 4, 0, 128, 0, 1)
	if err != nil {
		t.Error(err)
	}
	if w != 512 {
		t.Errorf("Run() width = %v, want = %v", w, 512)
	}
	if h != 341 {
		t.Errorf("Run() height = %v, want = %v", h, 341)
	}
	if svg == "" {
		t.Error("Run() svg is empty")
	}
}
