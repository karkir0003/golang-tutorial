package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**

	TODO: Write a program that takes in an integer. Make sure the integer is non-negative.

	Take the sum from 1 to that integer inclusive using a loop. Yes there's a math formula, but we want to practice loops here
	*/
	var input string
	fmt.Print("Enter a target integer: ")
	_, err := fmt.Scanln(&input) // Scan until the first space. If you want to scan the whole line, look into "bufio"

	// validate you got input
	if err != nil {
		fmt.Println("Error in receiving user input")
		return
	}

	// convert to integer attempt
	upperBoundSum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input must be an integer")
		return
	}

	// validate input
	if upperBoundSum < 0 {
		fmt.Println("Invalid input for upper bound sum. Enter a non-negative integer")
		return
	}
	sum := 0
	for i := 1; i <= upperBoundSum; i++ {
		sum += i
	}
	fmt.Println("Sum is: ", sum)
}
