package geometry

import "math"

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v1 Vector) Add(v2 Vector) (Vector) {
  return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vector) Subtract(v2 Vector) (Vector) {
  return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v1 Vector) DotProduct(v2 Vector) float64 {
  return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}
