//let I first start by writing a test case, then refactoring, debugging and make it to pass
package main

import (
	"testing"
)

// func TestTrial(t *testing.T) {
// 	got := hello("Yinebeb") //what the code gives us
// 	want := "hello Yinebeb" //the result which we are expecting to become.

// 	if want != got {
// 		t.Errorf("got%q,want%q", want, got)
// 	}
// }
//now let I add a subtest, inorder to share a common code:

func TestTrial(t *testing.T) {
	//@test   , subtest1  ; if test one passed but following subtest failed,
	//the compiler will report as it fiales. no way to show the pass of subtest1
	t.Run("Say hello to the people", func(t *testing.T) {
		got := hello("Tina", "fr") //what the code gives us
		want := "Bonjour,Tina"     //the result which we are expecting to become.
		// if got != want {
		// 	t.Errorf("got%q,want%q", got, want)
		// }
		assertmsg(t, got, want)
	})
	//@test , subtest2
	t.Run("Say hello to the world", func(t *testing.T) {
		got := hello("", "en")
		want := "hello World"
		// if got != want {
		// 	t.Errorf("got%q, want%q", got, want)
		// }
		assertmsg(t, got, want)
	})

	//@test, go on; let we refactor our code more; adding another requirement which take a languiage to and great asp.
	//@spanish
	t.Run("test spanish", func(t *testing.T) {
		got := hello("Tina", "sp")
		want := "Hola,Tina"
		assertmsg(t, got, want)
	})

	//@french
	t.Run("test French", func(t *testing.T) {
		got := hello("Tina", "fr")
		want := "Bonjour,Tina"
		assertmsg(t, got, want)
	})
	//@amh
	t.Run("test Amharic", func(t *testing.T) {
		got := hello("Tina", "amh")
		want := "Selam,Tina"
		assertmsg(t, got, want)
	})
}

// agian let I refactor the test an organized way- using a common assert_method
func assertmsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got%q,want%q", got, want)
	}
}
