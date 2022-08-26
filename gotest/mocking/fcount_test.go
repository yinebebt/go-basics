package mocking

import (
	"fmt"

	"github.com/cucumber/godog"
)

var nlist []int
var start int

func iStartCountingFrom(num int) error {
	start = num
	return nil
}

func iDecrementingBy(dec int) error {
	for i := start; i >= 0; i = i - dec {
		if i != 0 {
			printTheNumber(i)
		} else {
			ifTheNumebrIs(i)
		}

	}
	return nil
}

func ifTheNumebrIs(arg1 int) error {
	if arg1 == 0 {
		itShouldPrint("Go")
	}
	return nil
}

func itShouldPrint(str string) error {
	fmt.Println(str)
	return nil
}

func printTheNumber(num int) error {
	fmt.Println(num)
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I start counting from (\d+)$`, iStartCountingFrom)
	ctx.Step(`^I decrementing by (\d+)$`, iDecrementingBy)
	ctx.Step(`^Print (\d+)$`, printTheNumber)
	ctx.Step(`^if the numebr is (\d+)$`, ifTheNumebrIs)
	ctx.Step(`^it should print "([^"]*)"$`, itShouldPrint)
}
