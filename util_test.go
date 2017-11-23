package sqip

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"testing"
)

const testImage = "test/small.png"

func TestBase64(t *testing.T) {
	src := "some string"
	want := StringBase64("c29tZSBzdHJpbmc=")
	if got := Base64(src); got != want {
		t.Errorf("Base64() = %v, want %v", got, want)
	}
}

func TestImageTag(t *testing.T) {
	filename := "some/file/path.ext"
	svg := Base64(finalSVG)
	width := 1920
	height := 1275
	want := `<img width="1920" height="1275" src="some/file/path.ext" alt="Add descriptive alt text" style="background-size: cover; background-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxOTIwIDEyNzUiPjxmaWx0ZXIgaWQ9ImIiPjxmZUdhdXNzaWFuQmx1ciBzdGREZXZpYXRpb249IjEyIiAvPjwvZmlsdGVyPjxwYXRoIGQ9Ik0wIDBoMTkyMHYxMjc1SDB6IiBmaWxsPSIjZThkODAwIiAvPjxnIGZpbHRlcj0idXJsKCNiKSIgdHJhbnNmb3JtPSJzY2FsZSg3LjUwMDAwMCkgdHJhbnNsYXRlKDAuNSAwLjUpIj48ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgyMTYuNDE1NTg2IDY3LjA0MTg5OSkgcm90YXRlKDI2NC4xODUwMzIpIHNjYWxlKDEyLjU0NDgxMiAxMjIuODI0ODg4KSI+PGVsbGlwc2UgZmlsbD0iIzU5NDYwMCIgZmlsbC1vcGFjaXR5PSIuNTAxOTYxIiBjeD0iMCIgY3k9IjAiIHJ4PSIxIiByeT0iMSIgLz48L2c+PGVsbGlwc2UgZmlsbD0iIzdiNmQwMCIgZmlsbC1vcGFjaXR5PSIuNTAxOTYxIiBjeD0iMTA2IiBjeT0iNDUiIHJ4PSIxNiIgcnk9IjQxIiAvPjxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDQwIDExMCkgcm90YXRlKDIzOSkgc2NhbGUoMjQ1IDEzMCkiPjxwYXRoIGZpbGw9IiNmZjAiIGZpbGwtb3BhY2l0eT0iLjUwMTk2MSIgZD0iTS0uNS0uNWgxdjFoLTF6IiAvPjwvZz48ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgyNDguNzczNDQwIDU0LjQ2NzU3OSkgcm90YXRlKDE2OS45NDMwOTcpIHNjYWxlKDQ3Ljk2NDU4NyAyNC40MDQ3ODMpIj48ZWxsaXBzZSBmaWxsPSIjN2M2NTAwIiBmaWxsLW9wYWNpdHk9Ii41MDE5NjEiIGN4PSIwIiBjeT0iMCIgcng9IjEiIHJ5PSIxIiAvPjwvZz48cGF0aCBmaWxsPSIjY2ZiYTAwIiBmaWxsLW9wYWNpdHk9Ii41MDE5NjEiIGQ9Ik0xMjMuNDk4MjE2LTE2IDEwMC41OTI0OTQgNTguMjEwNTU3IDEwNC4yOTc5MDkgNzcuOTA1OTE2IDI1Ni4zNjk5MjIgNDkuMjI3ODExeiIgLz48ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgxMTcgNDkpIHJvdGF0ZSgxODkpIHNjYWxlKDY4IDIpIj48cGF0aCBmaWxsPSIjMDAwIiBmaWxsLW9wYWNpdHk9Ii41MDE5NjEiIGQ9Ik0tLjUtLjVoMXYxaC0xeiIgLz48L2c+PHBhdGggZmlsbD0iIzAwMCIgZmlsbC1vcGFjaXR5PSIuNTAxOTYxIiBkPSJNMjY1IDc2bC0xOS05LTM1IDExeiIgLz48ZWxsaXBzZSBmaWxsPSIjZmZmZDAwIiBmaWxsLW9wYWNpdHk9Ii41MDE5NjEiIGN4PSIxNTAiIGN5PSIxNTAiIHJ4PSIyNTUiIHJ5PSIyNiIgLz48L2c+PC9zdmc+);">`
	if got := ImageTag(filename, svg, width, height); got != want {
		t.Errorf("ImageTag() = %v, want %v", got, want)
	}
}

func Test_decode(t *testing.T) {
	src := "some string"
	want := StringBase64("c29tZSBzdHJpbmc=")
	if got := Base64(src); got != want {
		t.Errorf("Base64() = %v, want %v", got, want)
	}
}

func TestLoadImage(t *testing.T) {
	_, err := LoadImage(testImage)
	if err != nil {
		t.Error(err)
	}
}

func TestImageWidthAndHeight(t *testing.T) {
	img, err := LoadImage(testImage)
	if err != nil {
		t.Error(err)
	}
	w, h := ImageWidthAndHeight(img)
	if w != 512 {
		t.Errorf("ImageWidthAndHeight() width = %v, want %v", w, 512)
	}
	if h != 341 {
		t.Errorf("ImageWidthAndHeight() height = %v, want %v", h, 341)
	}
}

func Test_largerOne(t *testing.T) {
	if largerOne(1, 2) != 2 {
		t.Error("got 1, want 2")
	}
	if largerOne(2, 1) != 2 {
		t.Error("got 1, want 2")
	}
	if got := largerOne(1, 1); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}
