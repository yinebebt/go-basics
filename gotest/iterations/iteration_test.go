package iterations

import "testing"

func IterationTest(t *testing.T) {
	got := Repeated("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("got%s expected %s", got, expected)
	}

}
