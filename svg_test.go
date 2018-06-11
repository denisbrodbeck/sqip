package sqip

import (
	"testing"
)

const (
	minifiedSVG = `<svg xmlns="http://www.w3.org/2000/svg" width="1920" height="1275"><path d="M0 0h1920v1275H0z" fill="#e8d800" /><g transform="scale(7.500000) translate(0.5 0.5)"><g transform="translate(216.415586 67.041899) rotate(264.185032) scale(12.544812 122.824888)"><ellipse fill="#594600" fill-opacity=".501961" cx="0" cy="0" rx="1" ry="1" /></g><ellipse fill="#7b6d00" fill-opacity=".501961" cx="106" cy="45" rx="16" ry="41" /><g transform="translate(40 110) rotate(239) scale(245 130)"><path fill="#ff0" fill-opacity=".501961" d="M-.5-.5h1v1h-1z" /></g><g transform="translate(248.773440 54.467579) rotate(169.943097) scale(47.964587 24.404783)"><ellipse fill="#7c6500" fill-opacity=".501961" cx="0" cy="0" rx="1" ry="1" /></g><path fill="#cfba00" fill-opacity=".501961" d="M123.498216-16 100.592494 58.210557 104.297909 77.905916 256.369922 49.227811z" /><g transform="translate(117 49) rotate(189) scale(68 2)"><path fill="#000" fill-opacity=".501961" d="M-.5-.5h1v1h-1z" /></g><path fill="#000" fill-opacity=".501961" d="M265 76l-19-9-35 11z" /><ellipse fill="#fffd00" fill-opacity=".501961" cx="150" cy="150" rx="255" ry="26" /></g></svg>`
	finalSVG    = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1920 1275"><filter id="b"><feGaussianBlur stdDeviation="12" /></filter><path d="M0 0h1920v1275H0z" fill="#e8d800" /><g filter="url(#b)" transform="scale(7.500000) translate(0.5 0.5)"><g transform="translate(216.415586 67.041899) rotate(264.185032) scale(12.544812 122.824888)"><ellipse fill="#594600" fill-opacity=".501961" cx="0" cy="0" rx="1" ry="1" /></g><ellipse fill="#7b6d00" fill-opacity=".501961" cx="106" cy="45" rx="16" ry="41" /><g transform="translate(40 110) rotate(239) scale(245 130)"><path fill="#ff0" fill-opacity=".501961" d="M-.5-.5h1v1h-1z" /></g><g transform="translate(248.773440 54.467579) rotate(169.943097) scale(47.964587 24.404783)"><ellipse fill="#7c6500" fill-opacity=".501961" cx="0" cy="0" rx="1" ry="1" /></g><path fill="#cfba00" fill-opacity=".501961" d="M123.498216-16 100.592494 58.210557 104.297909 77.905916 256.369922 49.227811z" /><g transform="translate(117 49) rotate(189) scale(68 2)"><path fill="#000" fill-opacity=".501961" d="M-.5-.5h1v1h-1z" /></g><path fill="#000" fill-opacity=".501961" d="M265 76l-19-9-35 11z" /><ellipse fill="#fffd00" fill-opacity=".501961" cx="150" cy="150" rx="255" ry="26" /></g></svg>`
	// this is a problematic image - it's missing the default <g>-group
	specialSVGNoGroup = `<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 750 937'><path fill='#5e5145' d='M0 0h750v937H0z'/><path fill='#ffceae' fill-opacity='.5' d='M455.7 108l7.3 600.2L31.1 561.8z'/><ellipse fill-opacity='.5' rx='1' ry='1' transform='matrix(95.1777 101.02218 -375.52495 353.79954 302.3 76.1)'/><path fill='#fff1cf' fill-opacity='.5' d='M441 979l113.5-292.7L357 752.2z'/><ellipse fill='#00110b' fill-opacity='.5' rx='1' ry='1' transform='matrix(-118.08246 -24.86798 90.92995 -431.76942 634.4 845.2)'/><ellipse fill-opacity='.5' rx='1' ry='1' transform='matrix(-56.65586 45.16032 -103.38974 -129.70753 285.9 891.6)'/><path fill='#370000' fill-opacity='.5' d='M287.3 518h73.2v245h-73.2z'/><path fill='#fffff2' fill-opacity='.5' d='M130 543.5l164.6 172L258 499.7z'/><path fill='#baf8ff' fill-opacity='.5' d='M629 289.5l-2.5-73.2 124.3-4.3 2.6 73.1z'/><ellipse fill='#000002' fill-opacity='.5' rx='1' ry='1' transform='matrix(-125.36617 -5.50313 3.40138 -77.48636 645.6 36)'/><path fill='#001109' fill-opacity='.5' d='M382.6 434.2l33.7-60.8 134.5 74.5-33.8 60.8z'/></svg>`
	// problematic image with patch
	specialSVGPatched = `<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 750 937'><path fill='#5e5145' d='M0 0h750v937H0z'/><g filter='url(#c)' fill-opacity='.5'><path fill='#ffceae' fill-opacity='.5' d='M455.7 108l7.3 600.2L31.1 561.8z'/><ellipse fill-opacity='.5' rx='1' ry='1' transform='matrix(95.1777 101.02218 -375.52495 353.79954 302.3 76.1)'/><path fill='#fff1cf' fill-opacity='.5' d='M441 979l113.5-292.7L357 752.2z'/><ellipse fill='#00110b' fill-opacity='.5' rx='1' ry='1' transform='matrix(-118.08246 -24.86798 90.92995 -431.76942 634.4 845.2)'/><ellipse fill-opacity='.5' rx='1' ry='1' transform='matrix(-56.65586 45.16032 -103.38974 -129.70753 285.9 891.6)'/><path fill='#370000' fill-opacity='.5' d='M287.3 518h73.2v245h-73.2z'/><path fill='#fffff2' fill-opacity='.5' d='M130 543.5l164.6 172L258 499.7z'/><path fill='#baf8ff' fill-opacity='.5' d='M629 289.5l-2.5-73.2 124.3-4.3 2.6 73.1z'/><ellipse fill='#000002' fill-opacity='.5' rx='1' ry='1' transform='matrix(-125.36617 -5.50313 3.40138 -77.48636 645.6 36)'/><path fill='#001109' fill-opacity='.5' d='M382.6 434.2l33.7-60.8 134.5 74.5-33.8 60.8z'/></g></svg>`
	// final result of problematic image
	specialSVGFinal = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 750 937"><filter id="c"><feGaussianBlur stdDeviation="55" /></filter><path fill='#5e5145' d='M0 0h750v937H0z'/><g filter='url(#c)' fill-opacity='.5'><path fill='#ffceae' fill-opacity='.5' d='M455.7 108l7.3 600.2L31.1 561.8z'/><ellipse fill-opacity='.5' rx='1' ry='1' transform='matrix(95.1777 101.02218 -375.52495 353.79954 302.3 76.1)'/><path fill='#fff1cf' fill-opacity='.5' d='M441 979l113.5-292.7L357 752.2z'/><ellipse fill='#00110b' fill-opacity='.5' rx='1' ry='1' transform='matrix(-118.08246 -24.86798 90.92995 -431.76942 634.4 845.2)'/><ellipse fill-opacity='.5' rx='1' ry='1' transform='matrix(-56.65586 45.16032 -103.38974 -129.70753 285.9 891.6)'/><path fill='#370000' fill-opacity='.5' d='M287.3 518h73.2v245h-73.2z'/><path fill='#fffff2' fill-opacity='.5' d='M130 543.5l164.6 172L258 499.7z'/><path fill='#baf8ff' fill-opacity='.5' d='M629 289.5l-2.5-73.2 124.3-4.3 2.6 73.1z'/><ellipse fill='#000002' fill-opacity='.5' rx='1' ry='1' transform='matrix(-125.36617 -5.50313 3.40138 -77.48636 645.6 36)'/><path fill='#001109' fill-opacity='.5' d='M382.6 434.2l33.7-60.8 134.5 74.5-33.8 60.8z'/></g></svg>`
	// Test refit changes rect width
	preRefitSVG = `<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="189" height="280"><rect x="0" y="0" width="189" height="280" fill="#242221" /><g transform="scale(1.093750) translate(0.5 0.5)"><polygon fill="#ffffff" fill-opacity="0.501961" points="34,49 35,99 151,59" /></g></svg>`
	postRefitSVG = `<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="189" height="280"><rect x="0" y="0" width="190" height="280" fill="#242221" /><g transform="scale(1.093750) translate(0.5 0.5)"><polygon fill="#ffffff" fill-opacity="0.501961" points="34,49 35,99 151,59" /></g></svg>`
)

func Test_regexp(t *testing.T) {
	if matchSVGPath.MatchString(minifiedSVG) == false {
		t.Errorf("Regexp %q didn't match input.", matchSVGPath.String())
	}
	if matchSVGPath.MatchString(specialSVGNoGroup) == true {
		t.Errorf("Regexp %q matched input, but shouldn't.", matchSVGPath.String())
	}
	if matchPath.MatchString(minifiedSVG) == false {
		t.Errorf("Regexp %q didn't match test svg.", matchSVGPath.String())
	}
	if matchSVGClose.MatchString(minifiedSVG) == false {
		t.Errorf("Regexp %q didn't match test svg.", matchSVGPath.String())
	}
	if captureSVGOpen.MatchString(minifiedSVG) == false {
		t.Errorf("Regexp %q didn't match input.", captureSVGOpen.String())
	}
	matches := captureGroup.FindStringSubmatch(minifiedSVG)
	if matches == nil {
		t.Errorf("Regexp %q didn't capture input.", captureGroup.String())
	}
	matches = captureSVGOpen.FindStringSubmatch(minifiedSVG)
	if matches == nil {
		t.Errorf("Regexp %q didn't capture input.", captureSVGOpen.String())
	}
}

func Test_patchSVGGroup(t *testing.T) {
	want := specialSVGPatched
	got, err := patchSVGGroup(specialSVGNoGroup)
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("patchSVGGroup()\ngot = %v\nwant = %v", got, want)
	}
}

func TestBlur(t *testing.T) {
	want := finalSVG
	got, err := Blur(minifiedSVG, 1920, 1275)
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("Blur()\ngot = %v\nwant = %v", got, want)
	}
}

func TestBlur_special(t *testing.T) {
	want := specialSVGFinal
	got, err := Blur(specialSVGNoGroup, 750, 937)
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("Blur()\ngot = %v\nwant = %v", got, want)
	}
}

func TestRefit(t *testing.T) {
	want := postRefitSVG
	got := Refit(preRefitSVG, 190, 280)
	if got != want {
		t.Errorf("Refit()\ngot = %v\nwant = %v", got, want)
	}
}
