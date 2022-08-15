package main

import (
	"math"
)

// more on struct
type circle struct {
	x, y int
	r    float64
}

// passing variable-struct circle, by value.
func circleArea(c circle) float64 {
	return math.Pi * c.r * c.r
}

// passing variable-circle, by reference.
func circleArea_v1(c *circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *circle) circleArea_v2() float64 { //by passing a receiver btw func and funcName,
	//this will allows us to call a function using the . operator.
	//But if we can via dot, could it support definition of more than one method?
	return math.Pi * c.r * c.r

}

// per() int{
// 	return "null"
// }

// func main(){
// 	c:=circle{0,0,2.75}

// 	fmt.Println("Area of circel- V0: ", circleArea(c))
// 	fmt.Println("Area of circel- V1: ", circleArea_v1(&c))
// 	fmt.Println("Area of circel- V0: ", c.circleArea_v2())

// }
