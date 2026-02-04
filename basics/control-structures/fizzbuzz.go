package main

import "fmt"

func main() {
	/**
	Write a program that prints the numbers from 1 to 100,
	but for multiples of three, print “Fizz” instead of the number, and for the multiples of five, print “Buzz.”
	For numbers that are multiples of both three and five, print “FizzBuzz.”

	Also, allow the user to enter an integer
	*/

	currString := ""
	for i := 1; i <= 100; i++ {
		currString = ""
		if i%3 == 0 {
			currString += "Fizz"
		}
		if i%5 == 0 {
			currString += "Buzz"
		}
		if len(currString) > 0 {
			fmt.Println(currString)
		} else {
			fmt.Println(i)
		}
	}

	// Same solution but with while loop.
	// Go doesn't have a dedicated while keyword
	currString = ""
	ctr := 1
	fmt.Println("============ While Loop ===========")
	for ctr <= 100 {
		currString = ""
		if ctr%3 == 0 {
			currString += "Fizz"
		} else if ctr%5 == 0 {
			currString += "Buzz"
		}
		if len(currString) > 0 {
			fmt.Println(currString)
		} else {
			fmt.Println(ctr)
		}
		ctr += 1
	}
}
