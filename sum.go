package main

//returns the sum given an array of numbers
func Sum(nums []int) (sum int) {
	for _, val := range nums {
		sum += val
	}
	return
}

//returns the sum of slices as a single slice contraining the array sum in each value
func SumAll(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

//returns the sum of the tails of all slices as a single slice contraining the array sum in each value
func SumAllTails(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}
