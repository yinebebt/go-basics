/*Get started with golang, let we dig to go...*/

package main // declaring a package named 'main'.

import (
	"fmt"
	"unsafe"
)

// work with struct
type person struct {
	name    string
	age     int
	address string
}

func main() {
	fmt.Println("Hello, linux users this is golang project; july 27,2022@sheger tower")
	// practicing datatypes
	name := "AHello"
	aa := name[0] //this will proint the ascii code, 65 for A.
	fmt.Println(aa)

	// variable decalration
	var price int = 10 // default intialize to 0, string to ""
	salary := 1000
	var s string
	s = "Hello guys!"

	fmt.Println(price)
	fmt.Println(salary)
	fmt.Println(s)

	// array and slice: array is fixed sized, slice is variable size
	var a [5]int // array
	total := 0
	for i := range a {
		a[i] = i * 4
		total += a[i]
	}
	fmt.Println("average=", total/len(a))

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
	fmt.Println("Third value", count)

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

	bb := []int{4, 5, 6, 7, 8} // slice
	fmt.Println(bb)

	slice := make([]float64, 5, 10) // always slice is associated with an array,
	//as it is a segment of array.5 is the size of the slice, 10 is the max capacity the slice is pointed.
	arr := [10]float64{1, 2, 3, 4, 5, 6, 7, 0, 8, 9}
	slic2 := arr[2:5]
	fmt.Print("Slice:", slice, "\narr:", arr, "\nslice b/n 2:6", slic2)

	// loop- the only one is for loop:
	for i := 0; i < 3; i++ {
		fmt.Println("\n", i+2)
	}
	//to resembele the while loop:
	i := 0
	for i < 3 {
		fmt.Println(i + 2)
		i++
	}

	// panic and recover
	defer func() {
		str := recover() // str gets <nil>
		fmt.Println("Recovered from paic,", str)
	}()
	// panic("Panic - raised by programmer")

	//use struct: you can also declare struct as global.
	p := person{name: "Yinebeb", age: 12, address: "Addis"}
	fmt.Println("p is :", p)
	fmt.Println("p.age is::", p.age)

	type animal struct {
		spec   string
		age    int
		weight float64
	}
	//let we initialize the type-animal
	var ani animal                                    //or
	an := animal{spec: "cow", age: 12, weight: 45.89} // you can leave the colon and field name if you know the order.
	fmt.Println(ani, an)

}
