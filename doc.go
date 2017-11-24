// Package sqip allows SVG-based LQIP image creation
//
// https://github.com/denisbrodbeck/sqip
//
// https://godoc.org/github.com/denisbrodbeck/sqip/cmd/sqip
//
// This package is a go implementation of Tobias Baldauf‘s SVG-based LQIP technique
// (see https://github.com/technopagan/sqip).
//
// SQIP is an evolution of the classic LQIP technique: it makes use of Primitive to generate a SVG consisting of several simple shapes that approximate the main features visible inside the image, optimizes the SVG using minify and adds a Gaussian Blur filter to it.
// This produces a SVG placeholder which weighs in at only ~800-1000 bytes, looks smooth on all screens and provides an visual cue of image contents to come.
package sqip // import "github.com/denisbrodbeck/sqip"
