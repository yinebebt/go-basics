package main

import (
	"fmt"
	"goWithTest/testCase"
)

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(testCase.Morning)
}