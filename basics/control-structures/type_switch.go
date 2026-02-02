package main

import "fmt"

// GO BEST PRACTICE: Try to define type of your function! Go forces YOU to define the typing. Casting is done on
// a best effort basis
func processDay(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Unknown weekday")
	}
	return
}

func main() {
	/**
	TODO: Write a function that takes in a variable day.

	If day is an integer, then run a switch statement to decide what day of the week (only focus on weekdays)

	If day is another type, just print the type of the day
	*/

	var day any // generically typed variable with type determination done by us for concrete use case

	// pitfall: If you do day := 4, go compiler INFERS the type to ALWAYS be an int
	day = true

	// day.(type) is only meaningful for "any" type
	switch v := day.(type) {
	case int:
		processDay(v)
	default:
		fmt.Printf("Input type is: %T with value %v", v, v) // how to printf type of variable
	}
}
