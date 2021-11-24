package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func main() {

	icon, err := oksvg.ReadIcon("annular.svg", oksvg.WarnErrorMode)
	if err != nil {
		panic(err)
	}

	fmt.Printf("paths: %d\n", len(icon.SVGPaths))

	w, h := int(icon.ViewBox.W), int(icon.ViewBox.H)
	fmt.Printf("%dx%d\n", w, h)
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	scannerGV := rasterx.NewScannerGV(w, h, img, img.Bounds())
	raster := rasterx.NewDasher(w, h, scannerGV)
	icon.Draw(raster, 1.0)

	f, err := os.Create("annular.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
