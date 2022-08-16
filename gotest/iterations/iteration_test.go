package iterations

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Repeated("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("got '%s' expected '%s'", got, expected)
	}
}

// output's O need to be caps.
func ExampleRepeated() {
	str := Repeated("abc", 3)
	fmt.Println(str)
	//Output:abcabcabc
}

// let I add a benchmark
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeated("a", 3)
	}
}
