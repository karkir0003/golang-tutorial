package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	/**
	Parse user generated number, convert to float type and then assign the converted value to memory address of input variable
	*/
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println("Doubled number: ", output)
}
