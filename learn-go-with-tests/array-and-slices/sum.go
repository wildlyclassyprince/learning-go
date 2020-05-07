package main

// Sum adds elements of an array and returns the result
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
SumAll takes a varying number of slices and returns
a new slice containing the totals for each slice passed in.
*/
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
