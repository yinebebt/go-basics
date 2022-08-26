package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		countDown(buffer, &SpySleeper{})

		got := buffer.String()
		want := `3
2
1
Go`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpySleeper{}
		countDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			writemsg,
			sleepmsg,
			writemsg,
			sleepmsg,
			writemsg,
			sleepmsg,
			writemsg,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
