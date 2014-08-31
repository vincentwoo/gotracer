package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
	"github.com/vincentwoo/geometry"
)

func main() {
	img := renderImage(500, 500)

	outfile, _ := os.Create("out.png")
	defer outfile.Close()

	writer := io.Writer(outfile)

	png.Encode(writer, img)
}

func renderImage(width, height int) (*image.RGBA) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	eye   := geometry.Vector{-1, 0, 0}
	dir   := geometry.Vector{1, 0, 0}
	up    := geometry.Vector{0, 0.5, 0}
	down  := geometry.Vector{0, -0.5, 0}
	left  := geometry.Vector{0, 0, -0.5}
	right := geometry.Vector{0, 0, 0.5}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			xFactor := float64(x) / float64(width)
			yFactor := float64(y) / float64(height)

			leftComponent := left.Multiply(xFactor).Add(right.Multiply(1 - xFactor))
			upComponent   := up.Multiply(yFactor).Add(down.Multiply(1 - yFactor))

			img.Set(x, y, trace(eye, dir.Add(leftComponent).Add(upComponent).Normalize() ))
		}
	}

	return img
}

func trace(eye, dir geometry.Vector) color.Color {
	if math.Sqrt(dir.Y * dir.Y + dir.Z * dir.Z) < 0.3 {
		return color.RGBA{255, 0, 0, 255}
	}
	return color.RGBA{0, 255, 0, 255}
}
