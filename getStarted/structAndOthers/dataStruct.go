package structAndOthers

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
