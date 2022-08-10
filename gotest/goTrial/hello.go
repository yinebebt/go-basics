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
	// if lang == "sp" {
	// 	return spanish + "," + name
	// }
	// if lang == "fr" {
	// 	return french + "," + name
	// }
	// if lang == "amh" {
	// 	return amharic + "," + name
	// }
	// but instead of doing such lengthy if; let I use switch()
	switch lang {
	case "sp":
		return spanish + "," + name
	case "fr":
		return french + "," + name
	case "amh":
		return amharic + "," + name
	default:
		return def_eng + " " + name
	}

}
func main() {
	fmt.Println(hello("Yinebeb", "en"))

}
