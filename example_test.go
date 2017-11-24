package sqip_test

import (
	"log"
	"runtime"

	"github.com/denisbrodbeck/sqip"
)

func Example() {
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
