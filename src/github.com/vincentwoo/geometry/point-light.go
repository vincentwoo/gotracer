package geometry

import "image/color"

type PointLight struct {
	Origin   Vector
	Color    color.Color
	Strength float64
}
