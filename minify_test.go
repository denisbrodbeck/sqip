package sqip

import "testing"

const rawSVG = `<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="1920" height="1275">
<rect x="0" y="0" width="1920" height="1275" fill="#e8d800" />
<g transform="scale(7.500000) translate(0.5 0.5)">
<g transform="translate(216.415586 67.041899) rotate(264.185032) scale(12.544812 122.824888)"><ellipse fill="#594600" fill-opacity="0.501961" cx="0" cy="0" rx="1" ry="1" /></g>
<ellipse fill="#7b6d00" fill-opacity="0.501961" cx="106" cy="45" rx="16" ry="41" />
<g transform="translate(40 110) rotate(239) scale(245 130)"><rect fill="#ffff00" fill-opacity="0.501961" x="-0.5" y="-0.5" width="1" height="1" /></g>
<g transform="translate(248.773440 54.467579) rotate(169.943097) scale(47.964587 24.404783)"><ellipse fill="#7c6500" fill-opacity="0.501961" cx="0" cy="0" rx="1" ry="1" /></g>
<polygon fill="#cfba00" fill-opacity="0.501961" points="123.498216,-16.000000,100.592494,58.210557,104.297909,77.905916,256.369922,49.227811" />
<g transform="translate(117 49) rotate(189) scale(68 2)"><rect fill="#000000" fill-opacity="0.501961" x="-0.5" y="-0.5" width="1" height="1" /></g>
<polygon fill="#000000" fill-opacity="0.501961" points="265,76 246,67 211,78" />
<ellipse fill="#fffd00" fill-opacity="0.501961" cx="150" cy="150" rx="255" ry="26" />
</g>
</svg>`

func TestMinify(t *testing.T) {
	input := rawSVG
	output, err := Minify(input)
	if err != nil {
		t.Error(err)
	}
	// We can only test on shorter outputs, not on individual output comparison.
	// Package minify might make improvements to svg minification,
	// which in turn would change expected output.
	// Thus we only care, whether the minified output is shorter than the input.
	if len(output) >= len(input) {
		t.Errorf("Minify() result is not shorter than input.\noutput = %v\ninput = %v", output, input)
	}
}
