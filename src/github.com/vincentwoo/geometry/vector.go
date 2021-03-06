package geometry

import (
	"fmt"
	"image/color"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

// Conversion methods

func (v Vector) Color() color.Color {
	return color.RGBA64{
		uint16(math.Min(v.X*0xffff, 0xffff)),
		uint16(math.Min(v.Y*0xffff, 0xffff)),
		uint16(math.Min(v.Z*0xffff, 0xffff)),
		0xffff,
	}
}

// No-arg methods

func (v Vector) String() string {
	return fmt.Sprintf("(%.2f %.2f %.2f)", v.X, v.Y, v.Z)
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.DotProduct(v))
}

func (v Vector) Normalize() Vector {
	return v.Divide(v.Length())
}

// Takes another Vector

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v1 Vector) DotProduct(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vector) MultiplyV(v2 Vector) Vector {
	return Vector{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

// Takes a scalar

func (v1 Vector) Multiply(f float64) Vector {
	return Vector{v1.X * f, v1.Y * f, v1.Z * f}
}

func (v1 Vector) Divide(f float64) Vector {
	return Vector{v1.X / f, v1.Y / f, v1.Z / f}
}
