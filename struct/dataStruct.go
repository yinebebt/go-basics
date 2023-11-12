package main

import "fmt"

type Person struct {
	Name   string
	Height int
	Weight int
}

// let we embedd it with another struct; Runner is a person ; so he has a name and more
type Runner struct {
	Person // we passed all feature of person to Runner
	speed  float64
}

func (p Person) BMI() int {
	// fmt.Println("This gives us a BMI, from a struct")
	bmi_index := (p.Weight / p.Height)
	return bmi_index
}

func (r Runner) BMI_R() int {
	// fmt.Println("This gives us a BMI, from a struct")
	bmi_index := (r.Weight / r.Height)
	return bmi_index
}

func main() {
	//package struct
	P := Person{Name: "Yinebeb", Height: 2, Weight: 58}
	fmt.Printf("BMI Index of '%s' is %d", P.Name, P.BMI())

	R := Person{Name: "Haile", Height: 2, Weight: 62}
	fmt.Printf("BMI Index of '%s' is %d", R.Name, R.BMI())
}
