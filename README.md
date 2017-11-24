# SQIP does SVG-based LQIP image creation

<p align="center"><img src="logo.png" alt="Image of a blurred gopher" style="max-width:100%;"></p>

[![GoDoc](https://godoc.org/github.com/denisbrodbeck/sqip?status.svg)](https://godoc.org/github.com/denisbrodbeck/sqip) [![Go Report Card](https://goreportcard.com/badge/github.com/denisbrodbeck/sqip)](https://goreportcard.com/report/github.com/denisbrodbeck/sqip)

… because even blurred preview images need to look good :godmode:

## Overview

SQIP is a go implementation of [Tobias Baldauf's](https://tobias.is/) SVG-based LQIP [technique](https://github.com/technopagan/sqip).

[LQIP](http://www.guypo.com/introducing-lqip-low-quality-image-placeholders/) *(Low Quality Image Placeholders)* boils down to this:

* load the page initially with low quality images
* once the page loaded (e.g. in the onload event), replace them with full quality images

So instead of waiting for the final image to be rendered, we can serve a highly compressed image first, and then switch to the large one.

SQIP is an evolution of the classic LQIP technique: it makes use of [Primitive](https://github.com/fogleman/primitive) to generate a SVG consisting of several simple shapes that approximate the main features visible inside the image, optimizes the SVG using [minify](github.com/tdewolff/minify) and adds a Gaussian Blur filter to it.

This produces a SVG placeholder which weighs in at only *~800-1000* bytes, *looks smooth* on all screens and provides an *visual cue* of image contents to come.

## Installation

Get the cli app directly to your `$GOPATH/bin` with

```bash
go get -u github.com/denisbrodbeck/sqip/cmd/sqip
```

Import the library with

```golang
import "github.com/denisbrodbeck/sqip"
```

## CLI usage

> sqip [-n <int>] [-o <path>] [options...] <file>
> Flags:
>   -n  <int>     number of primitive SVG shapes (default: 8)
>   -o  <path>    save the placeholder SVG to a file (default: empty)
> Options:
>   -mode  <int>  shape type (default: 0)
>   -alpha <int>  color alpha (use 0 to let the algorithm choose alpha for each shape) (default: 128)
>   -bg    <hex>  background color as hex (default: avg)

```bash
# Generate a SVG placeholder and print an example <img> tag to stdout
sqip input.png

# Save the placeholder SVG to a file instead of printing the <img> to stdout
sqip -o output.svg input.png

# Customize the number of primitive SVG shapes (default=8) to influence bytesize or level of detail
sqip -n 4 input.jpg
```

## API

Here is an example app:

```golang
package main

import (
	"log"
	"runtime"
	"github.com/denisbrodbeck/sqip"
)

func main() {
	in := "path/to/image.png"   // input file
	out := "path/to/image.svg"  // output file
	workSize := 256             // large images get resized to this - higher size grants no boons
	count := 8                  // number of primitive SVG shapes
	mode := 1                   // shape type
	alpha := 128                // alpha value
	repeat := 0                 // add N extra shapes each iteration with reduced search (mostly good for beziers)
	workers := runtime.NumCPU() // number of parallel workers
	background := ""            // background color (hex)

  // create a primitive svg
	svg, width, height, err := sqip.Run(in, workSize, count, mode, alpha, repeat, workers, background)
	if err != nil {
		log.Fatal(err)
	}
	// save svg to file
	if err := sqip.SaveFile(out, svg); err != nil {
		log.Fatal(err)
	}
	// create example img tag
	tag := sqip.ImageTag(out, sqip.Base64(svg), width, height)
	log.Print(tag)
}
```

## Credits

* [Michael Fogleman](https://github.com/fogleman) and his awesome [primitive](https://github.com/fogleman/primitive) project
* [José Manuel Pérez](https://jmperezperez.com/about-me/) for his [awesome explanation](https://jmperezperez.com/svg-placeholders/) of [various](https://jmperezperez.com/more-progressive-image-loading/) [image](https://jmperezperez.com/lazy-loading-images/) [loading](https://jmperezperez.com/webp-placeholder-images/) [techniques](https://jmperezperez.com/medium-image-progressive-loading-placeholder/)
* [sqip](https://github.com/technopagan/sqip) for the research and initial nodejs implementation

The Go gopher was created by [Denis Brodbeck](https://github.com/denisbrodbeck), based on original artwork from [Renee French](http://reneefrench.blogspot.com/) and [Takuya Ueda](https://github.com/golang-samples/gopher-vector).

## License

The MIT License (MIT) — [Denis Brodbeck](https://github.com/denisbrodbeck). Please have a look at the [LICENSE](LICENSE) for more details.
