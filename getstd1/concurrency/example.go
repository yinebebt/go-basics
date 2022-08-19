// // example on concurrency

package main

import (
	"fmt"
	"time"
)

func hello(name string) {
	fmt.Println("hello routine started.")
	fmt.Println("hello ", name)
	// if I add this extra blocking, this hello routine wait for a millisecond and the scheduler will get chance to schedule
	// routines, the only one and active is hello and will take the race
	time.Sleep(10 * time.Millisecond)
	fmt.Println("hello routine ended.")

}

func main() {
	fmt.Println("Main execution started.")
	var input string
	//creating a goroutine
	go hello("Abele")
	//scheduling go routine, here goruntime will swap and execute the hello routine, sleeping the main,
	//from the multiplexing scheme
	fmt.Scanln(&input) // this is used as a blocking, to schedule goroutine; you can use time.sleep(10*Milliseconds too)

	fmt.Println("Main execution stopped.")
}
