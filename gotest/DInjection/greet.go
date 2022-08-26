package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func greet(writer io.Writer, name string) {

	// fmt.Printf("Hello %s", name)          // this will print to the stdout- the default console/terminal
	fmt.Fprintf(writer, "Hello %s", name) // this will print to the buffer/general purpose io.writer.
}

func greetHandler(h http.ResponseWriter, r *http.Request) {
	greet(h, "Written to http writer")
}

func main() {
	greet(os.Stdout, "Eyob")
	log.Fatal(http.ListenAndServe("127.0.0.1:5050", http.HandlerFunc(greetHandler)))
}
