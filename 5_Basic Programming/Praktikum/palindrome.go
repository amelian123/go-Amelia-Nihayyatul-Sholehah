package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Apakah Palindrome ? ")
	fmt.Scanln(&input)

	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
	if isPalindrome(input) {
		fmt.Printf("%s adalah palindrome\n", input)
	} else {
		fmt.Printf("%s adalah bukan palindrome\n", input)
	}
}
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
