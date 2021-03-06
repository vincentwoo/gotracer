package geometry

import "math"

type Sphere struct {
	Origin Vector
	Radius float64
	Material
}

func (sphere Sphere) Intersects(ray Ray) (intersects bool, intersection, normal Vector) {
	delta := ray.Origin.Subtract(sphere.Origin)

	b := 2 * ray.Direction.DotProduct(delta)
	c := delta.DotProduct(delta) - (sphere.Radius * sphere.Radius)

	discriminant := b*b - 4*c

	if discriminant <= 0 {
		return false, intersection, normal
	}

	t := (-b - math.Sqrt(discriminant)) / 2

	if t <= 0 {
		return false, intersection, normal
	}

	intersection = ray.Origin.Add(ray.Direction.Multiply(t))
	normal = intersection.Subtract(sphere.Origin).Divide(sphere.Radius)

	return true, intersection, normal
}
