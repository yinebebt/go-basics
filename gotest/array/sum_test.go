package array

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collecton of num-array", func(t *testing.T) {
		num := [5]int{1, 2, 3, 4, 5}
		got := Sum(num)
		want := 15

		if got != want {
			t.Errorf("got '%d' want '%d' %v", got, want, num) // the third arg is format string(?)
		}
	})
	t.Run("Collection of number-slices", func(t *testing.T) {
		num := []int{1, 2, 3, 4}
		got := SumSlice(num)
		want := 10

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
}

func ExampleSumSlice() {
	a := []int{4, 5, 6, 1}
	sum := SumSlice(a)
	fmt.Println(sum)
	//Output:16
}

func BenchmarkSum(b *testing.B) {
	a := [5]int{4, 5, 6, 5, 1}
	sum := Sum(a)
	fmt.Println("sum", sum)
}

// func sumAll
func TestSumAll(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}
	got := SumAll(a, b)
	want := []int{6, 22}

	for i, _ := range got {
		if got[i] != want[i] {
			t.Errorf("got '%d' want '%d' ", got, want)
		} else {
		}
	}
	//or you can use refelct.DeepEqual to check, but it is not type safe and needs care.
}

// func sumtail all
func TestAllTail(t *testing.T) {
	o := []int{} //empty slice, to remove runtime error,let we fix it
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}

	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		for i, _ := range got {
			if got[i] != want[i] {
				t.Errorf("got '%d' want '%d' ", got, want)
			}
		}
	}

	t.Run("Passing empty slice safely", func(t *testing.T) {
		got := SumTail(o, a)
		want := []int{0, 5}
		checkSum(t, got, want)
	})

	t.Run("Passing normal valued aslices", func(t *testing.T) {
		got := SumTail(a, b)
		want := []int{5, 18}
		checkSum(t, got, want)
	})

}
