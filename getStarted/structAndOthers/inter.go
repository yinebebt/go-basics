package structAndOthers

import (
	"math"
)

type shape interface {
	area() float64
	peremeter() float64
}

// let we create struct who implement the interface shape
type Circle struct {
	Rad float64
}
type Rect struct {
	Len, Wid float64
}

//below these struct will implement the interface by consuming all the methods shape have.

func (c Circle) area() float64 {
	return math.Pi * c.Rad * c.Rad
}

func (c Circle) peremeter() float64 {
	return 2 * math.Pi * c.Rad
}

func (r Rect) area() float64 {
	return r.Len * r.Wid
}

func (r Rect) peremeter() float64 {
	return 2 * (r.Len + r.Wid)
}

func (r Rect) newPeremeter() float64 { //here rect r act as a receiver, and possible to access this func as r.newPermeter
	return 2 * (r.Len + r.Wid)
}

func TotalArea(Shapes ...shape) float64 {
	var area float64
	for _, s := range Shapes {
		area += s.area()
	}
	return area
}

// func main() {
// 	totArea := totalArea(circle{4.5}, rect{4, 5})
// 	fmt.Println("Total Area is ", totArea)
// }
