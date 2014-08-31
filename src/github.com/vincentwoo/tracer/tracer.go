package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"github.com/vincentwoo/geometry"
)

func main() {
	img := renderImage(100, 100)

	outfile, _ := os.Create("out.png")
	defer outfile.Close()

	writer := io.Writer(outfile)

	png.Encode(writer, img)
}

func renderImage(width, height int) (*image.RGBA) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	eye := geometry.Vector{-5, 0, 0}
	dir := geometry.Vector{1, 0, 0}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, trace(eye, dir))
		}
	}

	return img
}

func trace(eye, dir geometry.Vector) color.Color {
	return color.RGBA{uint8(eye.X), 0, 0, 255}
}
