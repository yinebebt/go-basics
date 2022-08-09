package main

import (
	"fmt"
)

//once I passed the test, let I refactor by adding a constant and test agian.
const helloPrefix = "hello"

func hello(name string) string {
	return helloPrefix + " " + name
}
func main() {
	fmt.Println(hello("Yinebeb"))

}
