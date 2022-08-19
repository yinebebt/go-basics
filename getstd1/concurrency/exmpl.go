package main

import (
	"fmt"
	"time"
)

func getchar(str string) {
	for _, c := range str {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
	}
}

func getdigit(dig []int) {
	for _, d := range dig {
		fmt.Printf("%d at time %v\n", d, time.Since(start))
	}
}
