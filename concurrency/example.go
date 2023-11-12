// example on concurrency

package main

import (
	"fmt"
	"time"
)

var start time.Time
var a int

// init is go's special func which will initialize first
func init() {
	start = time.Now()
	a = 4
}

func main() {
	c := make(chan string)
	d := make(chan int)
	e := make(chan int, 3)

	fmt.Printf("Main execution started %v \n", time.Since(start))
	//var input string
	//creating a goroutine
	go hello("Abele")
	//scheduling go routine, here runtime will swap and execute the hello routine, sleeping the main,
	//from the multiplexing scheme
	//fmt.Scanln(&input) // this is used as a blocking, to schedule goroutine; you can use time.sleep(10*Milliseconds too)

	//at this level, go routines found below will not be considered, does it run at compile time?
	go getChar("hello")

	go getDigit([]int{1, 2, 3, 4, 5})
	go channel(c) //if channel comes her, things work fine, but if this is at the bottom, goroutines will fail.

	//anonymous go-routine
	go func() {
		fmt.Println("from anonymous go routine")
	}()
	// here both function definition and declaration goes in one

	// time.Sleep(10 * time.Millisecond)
	// go channel(c)   //here things will fail

	c <- "John"
	//	close(c) // closing the channel, belows line will cause a panic
	//	c <- "Mike"

	go square(d)
	//for loop, to block/unblock the main routine
	for val := range d { //you can use val,ok followed by if else ways
		fmt.Println(val)
		fmt.Println("from main routine")
	}
	go sum(e)

	e <- 10
	e <- 11
	e <- 12
	e <- 13 //here since channel comes full, blocking occur and go's runtime scheduler will invoked,routine will change

	fmt.Printf("Main execution stopped, a: %d\n", a)
}
