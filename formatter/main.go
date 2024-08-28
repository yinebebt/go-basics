package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := []byte("Hello")
	_ = "Hello" //string ia an alise for slice of byte.
	// when we assign new value, it will decode using utf8.

	fmt.Printf("%s\n", data)  // output: Hello, it
	fmt.Println(string(data)) // output: Helllo, it decode based on utf8 encoding
	fmt.Printf("%b\n", data)  // output: [1001000 1100101 1101100 1101100 1101111]
	fmt.Println(data)         // output: [72 101 108 108 111], default formatter - deciaml value of the byte
	fmt.Printf("%x\n", data)  //output: 48656c6c6f, hexadecimal

	// slice vs array
	// array are inflexible and type safae based on size
	// array, which is a slice with predefined size
	arrA := [3]int{}
	arrB := [4]int{}
	sliceA := []int{}
	sliceB := []int{}

	fmt.Println(reflect.TypeOf(arrA) == reflect.TypeOf(sliceA))   // false
	fmt.Println(reflect.TypeOf(arrA) == reflect.TypeOf(arrB))     // false
	fmt.Println(reflect.TypeOf(sliceB) == reflect.TypeOf(sliceB)) // true

	// cast array into slice
	sliceC := arrA[:]
	fmt.Println(reflect.TypeOf(arrA))
	fmt.Println(reflect.TypeOf(sliceC))
}
