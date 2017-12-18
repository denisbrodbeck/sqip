/*
Package main provides the command line app for SVG-based LQIP image creation.

Usage: sqip [-n <int>] [-o <path>] [options...] <file>

Flags:
  -n        <int>     number of primitive SVG shapes (default: 8)
  -o        <path>    save the placeholder SVG to a file (default: empty)
  -help     <bool>    show this help and exit
  -version  <bool>    show app version and exit
Options:
  -mode  <int>  shape type (default: 0)
  -alpha <int>  color alpha (use 0 to let the algorithm choose alpha for each shape) (default: 128)
  -bg    <hex>  background color as hex (default: avg)

If no output path is provided, an example <img> tag will be printed to stdout.

Available shape types:
 0=combo
 1=triangle
 2=rect
 3=ellipse
 4=circle
 5=rotatedrect
 6=beziers
 7=rotatedellipse
 8=polygon

Try:
  sqip -n 12 path/to/image.jpg
  sqip -mode 8 -o ./image.svg path/to/image.jpg
*/
package main // import "github.com/denisbrodbeck/sqip/cmd/sqip"

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/denisbrodbeck/sqip"
)

var (
	successExitCode    = 0
	errorExitCode      = 1
	errorParseExitCode = 2
)

func usage() {
	log.Println(usageStr)
	log.Println("Version:", version)
	os.Exit(errorParseExitCode)
}

func failOnErr(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(errorExitCode)
	}
}

var version = "master" // set by ldflags

func main() {
	var outFile string
	var count int
	var help bool
	var showVersion bool
	var mode int
	var alpha int
	var background string
	flag.StringVar(&outFile, "o", "", "")
	flag.StringVar(&outFile, "out", "", "")
	flag.IntVar(&count, "n", 8, "")
	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "help", false, "")
	flag.BoolVar(&showVersion, "version", false, "")
	flag.IntVar(&mode, "mode", 0, "")
	flag.IntVar(&alpha, "alpha", 128, "")
	flag.StringVar(&background, "bg", "", "")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if help {
		usage()
	}
	if showVersion {
		log.Println(version)
		os.Exit(successExitCode)
	}

	if flag.NArg() != 1 {
		log.Println("Missing input file")
		usage()
	}

	inFile := flag.Arg(0)
	workers := runtime.NumCPU()
	workSize := 256
	repeat := 0

	// seed random number generator for primitive
	rand.Seed(time.Now().UTC().UnixNano())

	svg, w, h, err := sqip.Run(inFile, workSize, count, mode, alpha, repeat, workers, background)
	failOnErr(err)

	if outFile == "" {
		fmt.Println(sqip.ImageTag(outFile, sqip.Base64(svg), w, h))
	} else {
		err = sqip.SaveFile(outFile, svg)
		failOnErr(err)
	}
}

const usageStr = `sqip is a tool for SVG-based LQIP image creation

Usage: sqip [-n <int>] [-o <path>] [options...] <file>
Flags:
  -n        <int>     number of primitive SVG shapes (default: 8)
  -o        <path>    save the placeholder SVG to a file (default: empty)
  -help     <bool>    show this help and exit
  -version  <bool>    show app version and exit
Options:
  -mode  <int>  shape type (default: 0)
  -alpha <int>  color alpha (use 0 to let the algorithm choose alpha for each shape) (default: 128)
  -bg    <hex>  background color as hex (default: avg)

If no output path is provided, an example <img> tag will be printed to stdout.
Available shape types: 0=combo 1=triangle 2=rect 3=ellipse 4=circle 5=rotatedrect 6=beziers 7=rotatedellipse 8=polygon

Try:
sqip -n 12 path/to/image.jpg
sqip -mode 8 -o ./image.svg path/to/image.jpg`
