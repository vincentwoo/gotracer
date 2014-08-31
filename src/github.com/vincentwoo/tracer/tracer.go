package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	// "math"
	"github.com/vincentwoo/geometry"
	"os"
)

func main() {
	img := renderImage(500, 500)

	outfile, _ := os.Create("out.png")
	defer outfile.Close()

	writer := io.Writer(outfile)

	png.Encode(writer, img)
}

func renderImage(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	eye := geometry.Vector{-1, 0, 0}
	dir := geometry.Vector{1, 0, 0}
	up := geometry.Vector{0, 0.5, 0}
	down := geometry.Vector{0, -0.5, 0}
	left := geometry.Vector{0, 0, -0.5}
	right := geometry.Vector{0, 0, 0.5}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			xFactor := float64(x) / float64(width)
			yFactor := float64(y) / float64(height)

			leftComponent := left.Multiply(xFactor).Add(right.Multiply(1 - xFactor))
			upComponent := up.Multiply(yFactor).Add(down.Multiply(1 - yFactor))

			color := trace(geometry.Ray{eye, dir.Add(leftComponent).Add(upComponent).Normalize()})
			img.Set(x, y, color)
		}
	}

	return img
}

func trace(ray geometry.Ray) color.Color {

	geo := geometry.Sphere{geometry.Vector{0, 0, 0}, 0.4, color.RGBA{255, 0, 0, 255}}
	light := geometry.DirectionalLight{geometry.Vector{-0.5, -1, 0}.Normalize(), color.RGBA{255, 255, 255, 255}, 1}

	if intersects, _, normal := geo.Intersects(ray); intersects {
		if f := normal.DotProduct(light.Direction); f > 0 {
			return color.RGBA{uint8(f * 255), 0, 0, 255}
		}
		return color.RGBA{0, 0, 0, 255}
	}
	return color.RGBA{150, 150, 150, 255}
}
