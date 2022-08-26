package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var numlist []int
var b *bytes.Buffer

func count(num, dec int) {
	for i := num; i > 0; i-- {
		if i != 0 {
			fmt.Fprintf(b, "%v", i) //since Println is not good for testing
		} else {
			fmt.Fprintf(b, "%v", "go")
		}
	}
}

//In BDD's feature step definition, since it is like testing;most of implementation need to focus on assertion.
// just practice along that.

//Mocking more ...

const writemsg = "write"
const sleepmsg = "sleep"

type sleeper interface {
	sleep()
}
type SpySleeper struct {
	Calls []string
}

func (s *SpySleeper) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, writemsg)
	return
} //here SpySleeper implement both sleeper interafe and io.writer; since it use both Write and sleep.
//NB: If you make write, it will be wrong.

func (s *SpySleeper) sleep() { // creating custom sleep func, since time.sleep staied longer.
	s.Calls = append(s.Calls, sleepmsg)
}

func countDown(out io.Writer, sleeper sleeper) { //here general purpose writer used
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.sleep()
	}
	fmt.Fprint(out, "Go")
}

func main() {
	sleeper := &SpySleeper{}

	countDown(os.Stdout, sleeper) //this will send to the console
}
