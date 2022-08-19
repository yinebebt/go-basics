package main

import "fmt"

func channel(name chan string) {
	var a chan int // empty/nil channel which has no any use,
	fmt.Println("value of empty channnel:", a)
	mychan := make(chan int) // this chan, formed via make comes useful to pass data

	fmt.Println("value of make channel: ", mychan)

	fmt.Println("hello" + <-name) // passing data from channel to stdout.

}
