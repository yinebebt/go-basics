package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{} //The Buffer type from the bytes package implements the Writer interface,meaning it can used as a palce to write sth.
	greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
