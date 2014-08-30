package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

type Vector struct {
	x, y, z float32
}

func main() {
	img := renderImage(100, 100)

	outfile, _ := os.Create("out.png")
	defer outfile.Close()

	writer := io.Writer(outfile)

	png.Encode(writer, img)
}

func renderImage(width, height int) (img *image.RGBA) {
	img = image.NewRGBA(image.Rect(0, 0, width, height))

	eye := Vector{-5, 0, 0}
	dir := Vector{1, 0, 0}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, trace(eye, dir))
		}
	}

	return
}

func trace(eye, dir Vector) color.Color {
	return color.RGBA{uint8(eye.x), 0, 0, 255}
}
