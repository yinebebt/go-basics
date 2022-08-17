package array

func Sum(a [5]int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func SumSlice(a []int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum

}

func SumAll(a ...[]int) []int {
	var allsum []int

	for _, slice := range a {
		sum := 0
		sum = SumSlice(slice)
		allsum = append(allsum, sum)
	}
	return allsum
}

func SumTail(a ...[]int) []int {
	var allsum []int

	for _, slice := range a {
		sum := 0
		num := len(slice)
		// for i, ele := range slice {
		// 	if i != 0 {
		// 		sum += ele
		// 	}
		//or

		if num != 0 {
			tail := slice[1:]
			sum = SumSlice(tail)
		} else {
			sum = 0
		}
		allsum = append(allsum, sum)
	}
	return allsum
}
