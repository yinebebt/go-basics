// example on concurrency
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(n int) {
	for i := 1; i < n; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(400)) // to see the gap
		time.Sleep(time.Millisecond * amt)
	}
	fmt.Println("")
}

func main() {
	//everything inside main will go first, sure?
	fmt.Println("Hello from main - goroutine")
	for i := 1; i < 4; i++ {
		go sum(3)          // invoking 4 routines
		fmt.Println("hhh") //why the routine skip this line
	}
	fmt.Println("hey, sum goroutine")
	var input string
	fmt.Scanln(&input)
}
