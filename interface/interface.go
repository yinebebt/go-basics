package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Peremeter() float64
}

// Circle : let we create struct who implement the interface shape
type Circle struct {
	Rad float64
}

type Rect struct {
	Len, Wid float64
}

// Area : below these struct will implement the interface by consuming all the methods Shape have.
func (c Circle) Area() float64 {
	return math.Pi * c.Rad * c.Rad
}

func (c Circle) Peremeter() float64 {
	return 2 * math.Pi * c.Rad
}

func (r Rect) Area() float64 {
	return r.Len * r.Wid
}

func (r Rect) Peremeter() float64 {
	return 2 * (r.Len + r.Wid)
}

// NewPeremeter : here rect r act as a receiver, and possible to access this func as r.NewPermeter
func (r Rect) NewPeremeter() float64 {
	return 2 * (r.Len + r.Wid)
}

func TotalArea(Shapes ...Shape) float64 {
	var area float64
	for _, s := range Shapes {
		area += s.Area()
	}
	return area
}

func main() {
	totArea := TotalArea(Circle{4.5}, Rect{4, 5})
	fmt.Println("Total area is ", totArea)
}
