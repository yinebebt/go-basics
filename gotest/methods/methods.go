package methods

import "math"

// main code...
type Rectangle struct {
	width, length float64
}
type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (r *Rectangle) Area() float64 { //?Area(rect Rectangle)
	return r.length * r.width
}

func (c *Circle) Area() float64 { //?cir Circle
	return math.Pi * c.radius * c.radius
}
