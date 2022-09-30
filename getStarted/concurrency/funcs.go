package main

import (
	"bytes"
	"fmt"
	"time"
)

func getchar(str string) {
	for _, c := range str {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
	}
}

func getdigit(dig []int) {
	for _, d := range dig {
		fmt.Printf("%d at time %v\n", d, time.Since(start))
	}
}

// exmples on pkg and functions
func pkg() {

	var buff bytes.Buffer // use this way to read/write to a bytes and strings.

	byt := buff.Bytes() // to convert tbuffer named buff into slice f bytes

	fmt.Println(byt)
}


func hello(name string) {

	fmt.Println("hello routine started.", time.Since(start))
	fmt.Println("hello ", name)
	// if I add this extra blocking, this hello routine wait for a millisecond and the scheduler will get chance to schedule
	// routines, the only one and active is hello and will take the race
	time.Sleep(10 * time.Millisecond)
	fmt.Println("hello routine ended.")

}