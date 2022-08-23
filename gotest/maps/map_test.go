package maps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("Defined word", func(t *testing.T) {
		dictionary := Dictionary{"Go": "GP programming languge"}
		got, _ := dictionary.search("Go")
		want := "GP programming languge"

		assertStrings(t, got, want)
	})

	t.Run("Undefined word", func(t *testing.T) {
		dictionary := Dictionary{"Go": "GP programming languge"}
		_, err := dictionary.search("python")
		want := errorNotFound
		if err == nil { //to ensure we don’t call .Error() on nil.
			t.Fatal("Expected error occured, may be  a nil pointer!") // if this happen, stop the test here
		}

		assertError(t, err, want) //err.Error() covert error into string
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{} //initialize empty map
		word := "C++"
		definition := "System programming language"
		err := dictionary.add(word, definition)

		assertError(t, err, nil) // compare the error with expected of normal state = nil
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing word", func(t *testing.T) {
		word := "C++"
		definition := "System programming language"
		dictionary := Dictionary{word: definition} // initialize with a pair, word:definition

		err := dictionary.add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.search(word)
	if err != nil {
		t.Fatal("First a word should added:", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "nginx" //engine-x
		definition := "fast and efficient web server"
		dictionary := Dictionary{word: definition}
		newDefinition := "mosty used server, next to apache"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "apache" //engine-x
		newDefinition := "mosty used server, next to apache"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrNoDeftoUpdate)
		// assertDefinition(t, dictionary, word, newDefinition)
	})
}


