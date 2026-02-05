package main

import (
	"fmt"
	"strings"
)

func joinStrings(separator string, s ...string) string {
	return strings.Join(s, separator) // join with separator
}

/*
*
Purview into strings library in Go
*/
func main() {
	// is target string contained as a substring in another string
	fmt.Println("ell substring of hello? ", strings.Contains("hello", "ell"))
	fmt.Println("eg substring of test? ", strings.Contains("test", "eg"))

	// count occurrences of smaller string in a bigger string
	fmt.Println("Number of t in attention? ", strings.Count("attention", "t"))

	// starts with
	fmt.Println("does alphabet start with alpha (prefix)? ", strings.HasPrefix("alphabet", "alpha"))

	// ends with
	fmt.Println("is bet a suffix of alphabet? ", strings.HasSuffix("alphabet", "bet"))

	// position of substring in string. -1 if not found
	fmt.Println("phab position in alphabet? ", strings.Index("alphabet", "phab"))

	// join array of strings
	fmt.Println(joinStrings(" ", "Alice", "Bob", "Charlie")) // space separator
	fmt.Println(strings.Join([]string{"a", "b", "c", "d", "e"}, "-"))

	// split string to array by character
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// to lower case
	fmt.Println(strings.ToLower("HaPpy BiRtHdAy"))

	// to upper case
	fmt.Println(strings.ToUpper("happy birthday"))
}
