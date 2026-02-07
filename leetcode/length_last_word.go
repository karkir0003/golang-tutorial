/*
*
Leetcode Problem: https://leetcode.com/problems/length-of-last-word/submissions/1911769299/?envType=problem-list-v2&envId=m5ao8cb6
*/
package main

import "strings"

func lengthOfLastWord(s string) int {
	cleanedString := strings.Trim(s, " ") // remove leading and trailing whitespace
	splitStrings := strings.Split(cleanedString, " ")

	lastWord := splitStrings[len(splitStrings)-1]
	return len(lastWord)
}
