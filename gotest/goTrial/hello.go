package main

import (
	"fmt"
)

//once I passed the test, let I refactor by adding a constant and test agian.
const helloPrefix = "hello"

const spanish = "Hola"
const french = "Bonjour"
const amharic = "Selam"

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "sp" {
		return spanish + "," + name
	}
	if lang == "fr" {
		return french + "," + name
	}
	if lang == "amh" {
		return amharic + "," + name
	}
	// but instead of doing such lengthy if; let I use switch()

	return helloPrefix + " " + name
}
func main() {
	fmt.Println(hello("Yinebeb", "en"))

}
