package main

import (
	"os"
	"time"

	"gotest/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
