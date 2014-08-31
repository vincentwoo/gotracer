package geometry

type Geometry interface {
  Intersects(r Ray) (bool, Vector, Vector)
}

