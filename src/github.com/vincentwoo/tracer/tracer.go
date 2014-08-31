package main

import (
	// "fmt"
	"github.com/vincentwoo/geometry"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
)

func main() {
	img := renderImage(500, 500)

	outfile, _ := os.Create("out.png")
	defer outfile.Close()

	writer := io.Writer(outfile)

	png.Encode(writer, img)
}

func renderImage(width, height int) *image.RGBA64 {
	img := image.NewRGBA64(image.Rect(0, 0, width, height))

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
	material := geometry.Material{
		Ambient:  geometry.Vector{29, 33, 38}.Divide(255),
		Diffuse:  geometry.Vector{34, 60, 150}.Divide(255),
		Specular: geometry.Vector{1, 1, 1},
	}
	geo := geometry.Sphere{
		Origin:   geometry.Vector{0, 0, 0},
		Radius:   0.4,
		Material: material,
	}
	light := geometry.DirectionalLight{
		Direction: geometry.Vector{-1, -1, 0}.Normalize(),
		Color:     geometry.Vector{1, 1, 1},
		Strength:  1,
	}

	if intersects, _, normal := geo.Intersects(ray); intersects {
		pixelColor := geo.Material.Ambient

		if f := normal.DotProduct(light.Direction); f > 0 {
			pixelColor = pixelColor.Add(
				light.Color.Multiply(f).MultiplyV(geo.Material.Diffuse))
		}

		reflectedV := light.Direction.Subtract(
			normal.Multiply(2 * light.Direction.DotProduct(normal)))

		reflected := reflectedV.DotProduct(light.Direction.Multiply(-1))

		if reflected > 0 {
			specular := math.Pow(reflected, 10)
			pixelColor = pixelColor.Add(
				light.Color.Multiply(specular).MultiplyV(geo.Material.Specular))
		}

		return pixelColor.Color()
	}
	return color.Transparent
}
