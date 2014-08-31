package geometry

import "image/color"

type DirectionalLight struct {
	Direction Vector
	Color     color.Color
	Strength  float64
}
