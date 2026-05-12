package main

import "slices"

func separateDigits(nums []int) []int {

	// Slice to store all separated digits
	digitsList := []int{}

	// Traverse the nums array from end to beginning
	for i := len(nums) - 1; i >= 0; i-- {

		// Get current number
		n := nums[i]

		// Extract digits one by one in reverse order
		// Example: 123 -> 3, 2, 1
		for n > 0 {

			// Get the last digit and append to result
			digitsList = append(digitsList, n%10)

			// Remove the last digit from the number
			n = n / 10
		}
	}

	// Reverse the slice to maintain original digit order
	slices.Reverse(digitsList)

	// Return the final separated digits array
	return digitsList
}

func main() {
	separateDigits([]int{123, 456, 789})
}