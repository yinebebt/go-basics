package main

import (
	"fmt"
)

//once I passed the test, let I refactor by adding a constant and test agian.

const spanish = "Hola"
const french = "Bonjour"
const amharic = "Selam"
const def_eng = "hello"

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greeting(lang) + "," + name
}

// again let I refactor the func...
func greeting(lang string) (prefix string) {
	switch lang {
	case "sp":
		prefix = spanish
	case "fr":
		prefix = french
	case "amh":
		prefix = amharic
	default:
		prefix = def_eng
	}
	return
}
func main() {
	fmt.Println(hello("Yinebeb", "en"))
}
