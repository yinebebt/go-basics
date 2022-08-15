package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// beyond a doc-comment; let I add an exmaple.
func ExampleAdd() { //
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6     // commenting this makes ExampleAdd to compile and execute;
	//but if you uncomment this,it will compile but not will execute.
}
