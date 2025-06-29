package leetcode

import "strconv"

// Given an integer x, return true if x is a palindrome, and false otherwise.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xAsString := strconv.Itoa(x)

	for i := range len(xAsString) / 2 {
		if xAsString[i] != xAsString[len(xAsString)-1-i] {
			return false
		}
	}

	return true

}
