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

	// first version
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
	fmt.Println("THird value", count)


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

	// a hash table - go's map
	m := make(map[int]string) // before use, map need to be initiated. this is iitiated wit pair 0,"" values
	m[1] = "Apple"
	m[2] = "Samsung"
	m[3] = "Nokia"
	for key, value := range m {
		fmt.Println("Key", key, "value", value)
	}
	delete(m, 2) // this remove the second element of map m.
	id, val := m[2]
	fmt.Println("id which is looked from the map at key 2:", id, "status either the lookup is successful or not:", val)
	//or
	if id1, ok := m[1]; ok { // if the lookup operation is successful,
		fmt.Println("from the if block of map", id1, ok)
	}

	//more on map:
	employee := map[string]map[string]float64{
		"Tigst": {
			"salary": 10000,
			"age":    24,
		},
		"Abel": {
			"experience": 4,
			"salary":     41000,
			"age":        36,
		},
		"Yina": map[string]float64{ // also possible to redecalre here
			"woreda": 12,
			"age":    23,
		},
	}
	if emp, ok := employee["Abel"]; ok {
		fmt.Println("Experience", emp["experience"], "Salary:", emp["salary"])
	}

	greet("yina")

	// call the sum func
	result := sum(salary, price)
	fmt.Println("result come from func-sum,", result)
	Rsqrt, messsage := sqrt(58)
	fmt.Println("From the SQRT func: ", Rsqrt, messsage)

	fmt.Println("RESULT Retruned from f1:", f1())

	fmt.Println("Sum of args is:", add(1, 2, 3))

	nextEven := evengenerator()
	fmt.Println("First Even:", nextEven())
	fmt.Println("Second Even:", nextEven())
	fmt.Println("Third Even:", nextEven())

	fmt.Println("Factorial of num:", fact(5))
	fmt.Println("Fibonacci of num:", fibon(4))
	defer sec()
	fir()

	// panic and recover
	defer func() {
		str := recover()
		fmt.Println("Recovered from paic!", str)
	}()
	panic("Panic")

	//use struct: you can also declare struct as global.
	p := person{name: "Yinebeb", age: 12, address: "Addis"}
	fmt.Println(p)
	fmt.Print(p.age)

	type animal struct {
		spec   string
		age    int
		weight float64
	}
	//let we initialize the type-animal
	var ani animal                                    //or
	an := animal{spec: "cow", age: 12, weight: 45.89} // you can leave the colon and field name if you know the order.
	fmt.Println(ani, an)
	type circle struct {
		xl, yl, rad float64
	}
	C := circle{0.0, 0.0, 2.57}
	fmt.Println(areaCircle(C.xl, C.yl, C.rad))
}
