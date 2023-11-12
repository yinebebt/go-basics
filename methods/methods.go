package main

import (
	"errors"
	"fmt"
	"math"
)

// function outside main

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined negative input")
	}
	return math.Sqrt(x), nil
}

// function call stack-trace
func f1() (r int) {
	r = f2()
	return
}

func f2() int {
	return 4
}

// variadic function
func add(arg ...int) int {
	total := 0
	for _, v := range arg {
		total += v
	}
	return total
}

func areaCircle(x, y int, r float64) float64 {
	return math.Pi * r * r
}

// a function that return another function-closure
// If the function is used in a return value or to be assigned to variable, no need to mention a name.
func evenGenerator() func() uint {
	i := uint(0)               // to make 0 unit-type casting from int to uint.
	return func() (ret uint) { // the first return value is 0.
		ret = i
		i += 2
		return
	}
}

// recursion func.
func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}

func fibon(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	}
	return fibon(num-1) + fibon(num-2)
}
func fir() {
	fmt.Println("first")
}
func sec() {
	fmt.Println("second")
}

// function with a pointer: to demonstrate that a value will copy appropriately.
func zero(xptr *int) {
	*xptr = 0
}

// more on pointers:
func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	temp := *y // temp initialized as a pointer and store value of y = *y
	*y = *x    // assign value of x to y
	*x = temp  // assign value stored on pointer temp which is y's initial value to pointer x.
	//*px, *py = *py, *px   //this approach seems by applying concurrency, go assign both value at a time;
	// the change will reflect after this line goes run, I think. But this is not the fact; rather it is due right-side
	// expressions will evaluate first, then assignment goes right to left.
}

// Refactoring long-ifs into a better way:function which greets name in a particular language
var greetings = map[string]string{ // globally declarations need a var keyword.
	"en":  "Hello",
	"fr":  "Bounjar",
	"it":  "itali",
	"or":  "Akam",
	"amh": "ሰለመ",
}

func greet(name string) {
	for id := range greetings {
		if welcome, ok := greetings[id]; ok {
			fmt.Println(welcome, name)
		} else {
			errors.New("language not found!")
		}
	}
}

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
func circleAreaV1(c *circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *circle) circleAreaV2() float64 {
	// by passing a receiver btw func and funcName, this will allow us to call a function using the . operator.
	// But if we can via dot, could it support definition of more than one method?
	return math.Pi * c.r * c.r
}
