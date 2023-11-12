/*
For more, checkout
	https://securego.io/docs/rules/g304.html
	https://github.com/securego/gosec
*/

package gosec

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// run `gosec -include=G101 ./gosec`
func gosecG101() {
	username := "admin"
	var password = "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"

	fmt.Println("Doing something with: ", username, password)
}

//don't use literals like password, token as a variable name.

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

// declare the url as constant instead of var
