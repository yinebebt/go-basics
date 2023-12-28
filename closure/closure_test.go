package main

import (
	"fmt"
	"testing"
)

func TestEvaluate(t *testing.T) {
	t.Run("No errors when x is less than 10", func(t *testing.T) {
		errs := evaluate(5)
		if len(errs) != 0 {
			t.Errorf("Expected no errors, got: %v", errs)
		}
	})

	t.Run("Errors when x is greater than 10", func(t *testing.T) {
		errs := evaluate(15)
		expected := []error{fmt.Errorf("15 is over 10\n")}
		if !compareErrors(errs, expected) {
			t.Errorf("Expected errors %v, got: %v", expected, errs)
		}
	})

	t.Run("Errors when x is greater than 20", func(t *testing.T) {
		errs := evaluate(25)
		expected := []error{fmt.Errorf("25 is over 10\n"), fmt.Errorf("25 is over 20\n")}
		if !compareErrors(errs, expected) {
			t.Errorf("Expected errors %v, got: %v", expected, errs)
		}
	})
}

func compareErrors(a, b []error) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Error() != b[i].Error() {
			return false
		}
	}
	return true
}
