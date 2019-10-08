package main

import "fmt"

// Sum of array or slice
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll sums of list of slices
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func main() {
	var sum = SumAll([]int{1, 2}, []int{0, 9})

	fmt.Println(sum)
}
