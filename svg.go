package sqip

import (
	"errors"
	"fmt"
	"regexp"
)

// regular minified svg
var matchSVGPath = regexp.MustCompile(`^<svg.*?><path.*?><g`)
var matchPath = regexp.MustCompile(`<path.*?>`)
var matchSVGClose = regexp.MustCompile(`</svg>`)

// capture <g only once (lazy)
var captureGroup = regexp.MustCompile(`^(.*?)(<g)(.*?)$`)
var captureSVGOpen = regexp.MustCompile(`(<svg)(.*?)(>)`)

//Capture the (pre-minified) background rectangle
var captureBackground = regexp.MustCompile(`<rect.*fill="(.*)".*/>`)

// Blur adds viewbox and preserveAspectRatio attributes as well as
// a Gaussian Blur filter to the SVG.
// When no group is found, add group (element with blur applied) using patchSVGGroup().
func Blur(inSVG string, width, height int) (outSVG string, err error) {
	blurStdDev := 12
	blurFilterID := "b"
	outSVG = inSVG
	if matchSVGPath.MatchString(inSVG) == false {
		blurStdDev = 55
		outSVG, err = patchSVGGroup(outSVG)
		if err != nil {
			return "", err
		}
		blurFilterID = "c"
	} else {
		matches := captureGroup.FindStringSubmatch(inSVG)
		if matches == nil {
			return "", errors.New("Failed to capture group")
		}
		outSVG = fmt.Sprintf("%s%s%s", matches[1], `<g filter="url(#b)"`, matches[3])
	}

	repl := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d"><filter id="%s"><feGaussianBlur stdDeviation="%d" /></filter>`, width, height, blurFilterID, blurStdDev)

	return captureSVGOpen.ReplaceAllString(outSVG, repl), nil
}

// Add Group to SVG.
// Use only on malformed svgs with no groups.
func patchSVGGroup(svg string) (string, error) {
	loc := matchPath.FindStringIndex(svg)
	if loc == nil {
		return "", errors.New("failed to find path in svg")
	}
	startIndex := loc[1] // end of matched string
	loc = matchSVGClose.FindStringIndex(svg)
	if loc == nil {
		return "", errors.New("failed to find svg close tag in svg")
	}
	endIndex := loc[0] // start of matched string

	group := `<g filter='url(#c)' fill-opacity='.5'>`

	res := fmt.Sprintf("%s%s%s</g></svg>", svg[0:startIndex], group, svg[startIndex:endIndex])
	return res, nil
}

// Refit adjusts the width of the background to exactly match the output size
// This will prevent rounding errors causing the aspect ratio to become incorrect
// This must be called before minify
func Refit(in string, width int, height int) (out string) {
	return captureBackground.ReplaceAllString(in, fmt.Sprintf("<rect x=\"0\" y=\"0\" width=\"%d\" height=\"%d\" fill=\"${1}\" />", width, height))
}
