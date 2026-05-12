package main

import "fmt"

func isPalindrome(x int) bool {

	// If the number is negative or ends in 0 (except for 0 itself), it is not a palindrome
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// Initialse Variable to store the reversed number
	reverse := 0

	// Create a copy of x so the original value remains unchanged
	xCopy := x

	// Reverse the digits of the number
	for xCopy > 0 {

		// Extract the last digit and append it to reverse
		reverse = (reverse * 10) + (xCopy % 10)

		// Remove the last digit from xCopy
		xCopy /= 10
	}

	// Check if the reversed number is equal to the original number
	if x == reverse {
		// If equal, return true
		return true
	}

	// If not equal, return false
	return false
}

func main() {
	fmt.Println(isPalindrome(121))
}
