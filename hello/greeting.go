// Refactoring long ifs into a better way:function which greets name in a particular language
package main

import (
	"errors"
	"fmt"
)

var greetings = map[string]string{ // globally dcalarations need a var keyword.
	"en":  "Hello",
	"fr":  "Bounjar",
	"it":  "itali",
	"or":  "Akam",
	"amh": "ሰለመ",
}

func greet(name string) {
	for id, _ := range greetings {
		if welcom, ok := greetings[id]; ok {
			fmt.Println(welcom, name)
		} else {
			errors.New("Language not found!")
		}
	}
}
