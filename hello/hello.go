/*Get started with golang, let we dig to go...*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello, linux users this is golang project; july 27,2022@sheger tower")
	// practicing datatypes
	name := "AHello"
	a := name[0] //this will proint the ascii code, 65 for A.
	fmt.Println(a)

	var b int = 21
	var c = "Saint Mary"
	fmt.Println(b)
	fmt.Printf("%d,%s", b, c)

	fmt.Println(len("ABC 123"))
	fmt.Println("Hello" + "﹵፪፩፯")
	fmt.Println(unsafe.Sizeof(name)) // this print size of strings descriptor, which is pointers size + lengths -wich is int, for 64-bit ssytem,this is 16byte.

	var (
		company      = "2f-capital"
		address      = "Bole, Sheger-tower"
		employeeSize = ">30"
	)
	const pi float64 = 3.14
	fmt.Printf("Company:%q,\nAddress:%s,\nEmpSize:%s\n", company, address, employeeSize)
	// work with inputs
	var inputAge int
	fmt.Print("ENter your age:")
	fmt.Scanf("%d", &inputAge) // if you take the value,Scanf will return two values, and a bit hard to gt the value. check it
	fmt.Println("Age entered:", inputAge)

	fmt.Println(`1    
2
3
 	`) //new line will taken as a semicolon in go's script.

	//increments, a bit subtle fro other languages
	var count = 5
	fmt.Println("First value", count)
	count++
	fmt.Println("second value", count)
	count--
	fmt.Println("THird value", count)
	//fmt.Println(count++)  // here increments are removed for expression, not as in other like c,c++, java.
	//You know as increment/decrement doesn't exist in python.
	// also prefix like --count also an error in go.

	//work with loop, conditions
	var odd []int
	var even []int
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			even = append(even, i)

		} else {
			odd = append(odd, i)
		}
	}
	fmt.Println("Even # between 1 and 10:", even)
	fmt.Print("or here in one:")
	for i := 1; i < len(even); i++ {
		fmt.Print(", ", even[i])
	}
	fmt.Println("\nOdd # between 1&10:", odd)
	fmt.Print("or here in one:")
	for i := 1; i < len(odd); i++ {
		fmt.Print(", ", odd[i])
	}

	var grade int
	for j := 1; j < 4; j++ {
		fmt.Print("\nEnter your exam grade: ")
		fmt.Scanf("%d", &grade)
		switch grade {

		case 40:
			fmt.Println("Your grade is fourty.")
		case 80:
			fmt.Println("Your grade is eighty.")
		case 100:
			fmt.Println("Wow, really you are smart")
			fmt.Println("we will hire you!")

		default:
			fmt.Println("NO grade given!")
		}
	}

	//more on go datastructures-array,slice, map
	var emp [4]string
	emp2 := [4]string{"Ezedin", "Senait", "Kerim", "Zeray"}
	emp3 := [3]string{
		"Ezedin",
		"Senait",
		//"Zeray",
		"Josef",
	}
	fmt.Println(emp)
	fmt.Println(emp2)
	fmt.Println(emp3)

	greet("yina")

}
