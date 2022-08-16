package main

import (
	"fmt"
)

func Hello(name, language string) string {
	return fmt.Sprintf("%s, %s", greeting(language), name)
}

var greetings = map[string]string{
	"es": "Hola",
	"fr": "Bonjour",
	//etc..
}

func greeting(language string) string { //breakPoint - just to get the idea...
	greeting, exists := greetings[language]
	if exists {
		return greeting
	}
	return "Hello"
}

func main() {
	result := Hello("Yinebeb", "fr")
	fmt.Println(result)

	//from goWithTest
	//fmt.Println(Hello())

}
