package geometry

import "math"

type Sphere struct {
  Origin Vector
  Radius float64
}

func (sphere Sphere) Intersects(ray Ray) bool {
  delta := ray.Origin.Subtract(sphere.Origin)

  // a := 1
  b := 2 * ray.Direction.DotProduct(delta)
  c := delta.DotProduct(delta) - (sphere.Radius * sphere.Radius)

  if c < 0 {
    return false
  }

  return -b - math.Sqrt(4 * c) > 0
}
