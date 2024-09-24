// Package struct demonstrates how to define and use structs in Go.
package main

import "fmt"

type Person struct {
	Name   string
	Height int
	Weight int
}

// Runner let we embed it with another struct;
// Runner is a person, so he has a name and more
type Runner struct {
	Person // we passed all feature of person to Runner
	speed  float64
}

// BMI returns a BMI, from the struct
func (p Person) BMI() int {
	bmiIndex := p.Weight / p.Height
	return bmiIndex
}

func (r Runner) BMIR() int {
	bmiIndex := r.Weight / r.Height
	return bmiIndex
}

func main() {
	P := Person{Name: "Yinebeb", Height: 2, Weight: 58}
	fmt.Printf("BMI Index of '%s' is %d", P.Name, P.BMI())

	R := Person{Name: "Haile", Height: 2, Weight: 62}
	fmt.Printf("BMI Index of '%s' is %d", R.Name, R.BMI())
}
