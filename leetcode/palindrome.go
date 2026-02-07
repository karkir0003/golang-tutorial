package main

import "fmt"

func palindrome(s string) bool {
	// check if string is a palindrome
	start := 0
	end := len(s) - 1
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}

func main() {
	s1 := "racecar"
	s2 := "aabbc"
	fmt.Println("s1 palindrome: ", palindrome(s1))
	fmt.Println("s2 palindrome: ", palindrome(s2))
}
