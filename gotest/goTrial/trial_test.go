//let I first start by writing a test case, then refactoring, debugging and make it to pass
package main

import (
	"testing"
)

func TestTrial(t *testing.T) {
	got := hello("Yinebeb") //what the code gives us
	want := "hello Yinebeb" //the result which we are expecting to become.

	if want != got {
		t.Errorf("got%q,want%q", want, got)
	}
}
