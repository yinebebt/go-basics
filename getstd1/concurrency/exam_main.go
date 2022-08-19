// // example on concurrency

package main

import (
	"fmt"
	"time"
)

var start time.Time
var a int

func init() { // this init, go's special func will initlaize first.
	start = time.Now()
	a = 4
}

func hello(name string) {

	fmt.Println("hello routine started.", time.Since(start))
	fmt.Println("hello ", name)
	// if I add this extra blocking, this hello routine wait for a millisecond and the scheduler will get chance to schedule
	// routines, the only one and active is hello and will take the race
	time.Sleep(10 * time.Millisecond)
	fmt.Println("hello routine ended.")

}

func main() {
	c := make(chan string)
	d := make(chan int)

	fmt.Printf("Main execution started %v \n", time.Since(start))
	//var input string
	//creating a goroutine
	go hello("Abele")
	//scheduling go routine, here goruntime will swap and execute the hello routine, sleeping the main,
	//from the multiplexing scheme
	//fmt.Scanln(&input) // this is used as a blocking, to schedule goroutine; you can use time.sleep(10*Milliseconds too)

	//at this level, go routines found below will not considered, doe it doen at compile time?
	go getchar("hello")

	go getdigit([]int{1, 2, 3, 4, 5})
	go channel(c) //if cannel comes her, things work fine, but if this is at the bottom, goroutines will failed.

	//anonymous go-routine
	go func() {
		fmt.Println("from anonymous go routine")
	}()
	// here both function definition and declaration goes in one

	// time.Sleep(10 * time.Millisecond)
	// go channel(c)   //here things will failed

	c <- "John"
	//	close(c) // closing the channel, belows line will cause a panic
	//	c <- "Mike"

	go square(d)
	//for loop, to block/unnblock the main routine
	for val := range d { //you can use val,ok followed by if else ways
		fmt.Println(val)
		fmt.Println("from main routine")
	}
	fmt.Printf("Main execution stopped, a: %d\n", a)
}
