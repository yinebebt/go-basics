//practice pointer on map, struct and slice here ...
// taking a structure
type Employee struct {

	// taking variables
	name  string
}

// Main Function
func main() {
	emp := Employee{"Abebe"}
	emp_add := &emp
	emp_add2 := &emp
	emp2 := newEmp()
	emp2.name:="kebede"

	fmt.Println(emp_add)
	fmt.Println(emp_add2)

	fmt.Println(emp2.name)
}

func newEmp() *Employee {
	return &Employee{string}
}