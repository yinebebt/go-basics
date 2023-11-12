package main

import "fmt"

func main() {
	// work with loop, conditions
	var odd []int
	var even []int
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			even = append(even, i)

		} else {
			odd = append(odd, i)
		}
	}
	fmt.Println("Even # between 1 and 10:", even)
	fmt.Print("or here in one:")
	for i := 1; i < len(even); i++ {
		fmt.Print(", ", even[i])
	}
	fmt.Println("\nOdd # between 1&10:", odd)
	fmt.Print("or here in one:")
	for i := 1; i < len(odd); i++ {
		fmt.Print(", ", odd[i])
	}

	var grade int
	for j := 1; j < 4; j++ {
		fmt.Print("\nEnter your exam grade: ")
		fmt.Scanf("%d", &grade)
		switch grade {

		case 40:
			fmt.Println("Your grade is forty.")
		case 80:
			fmt.Println("Your grade is eighty.")
		case 100:
			fmt.Println("Wow, really you are smart")
			fmt.Println("We will hire you!")

		default:
			fmt.Println("No grade given!")
		}
	}
}
