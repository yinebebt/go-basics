// Package methods demonstrates how to define and use methods in Go.
package main

import "fmt"

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is still 5, but if we use a pointer in func zero, it will be 0.
	y := 1.5
	square(&y) // after this y will be 2.25

	C := circle{0.0, 0.0, 2.57}
	fmt.Println(areaCircle(C.x, C.y, C.r)) //the float 0.0 will chunk-off to int 0 implicitly.

	greet("Yinebeb")

	// call the sum func
	result := sum(1000, 145)
	fmt.Println("result come from func-sum,", result)
	Sqrt, message := sqrt(58)
	fmt.Println("From the SQRT func: ", Sqrt, message)

	fmt.Println("result returned from f1:", f1())

	fmt.Println("Sum of args is:", add(1, 2, 3))

	nextEven := evenGenerator()
	fmt.Println("First Even:", nextEven())
	fmt.Println("Second Even:", nextEven())
	fmt.Println("Third Even:", nextEven())

	fmt.Println("Factorial of num:", fact(5))
	fmt.Println("Fibonacci of num:", fibon(4))
	defer sec()
	fir()

	c := circle{0, 1, 2.75}

	fmt.Println("Area of circle - V0: ", circleArea(c))
	fmt.Println("Area of circle - V1: ", circleAreaV1(&c))
	fmt.Println("Area of circle - V0: ", c.circleAreaV2())
}
