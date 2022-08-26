package mocking

import (
	"bytes"
	"fmt"
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
