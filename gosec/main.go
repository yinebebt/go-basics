// Package gosec demonstrates how to scan Go code for potential security vulnerabilities using Gosec.
// For more, checkout https://securego.io/docs/rules/g304.html, https://github.com/securego/gosec
package gosec

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// gosecG101 don't use literals like password, token as a variable name.
func gosecG101() {
	username := "admin"
	var password = "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"

	fmt.Println("Doing something with: ", username, password)
}

// gosecG107 declare the url as constant instead of var
func gosecG107() {
	var url string = "https://www.google.com"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}

// Note to run, use `$gosec -include=G101 ./gosec` command
