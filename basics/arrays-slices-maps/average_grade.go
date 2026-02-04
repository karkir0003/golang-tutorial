package main

import (
	"fmt"
)

func main() {
	// manually enter 5 grades for example

	var grades [5]float64

	grades[0] = 98
	grades[1] = 93
	grades[2] = 77
	grades[3] = 82
	grades[4] = 83

	total := float64(0) // cast integer to float

	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	avg := total / float64(len(grades)) // you have to manually cast int to float64 for len due to len function returning int

	fmt.Printf("Average grade is: %.2f\n", avg) // Round average to two decimal places

	// Enhanced for loop in go for the same
	// for index, value := range grades is like for i, val in enumerate(grades) in Python. You get convenient
	// access to the index and value in array
	total = float64(0)

	// _ means "undeclared var". Go doesn't allow for unused variables in compilation
	for _, grade := range grades {
		total += grade
	}

	avg = total / float64(len(grades))

	fmt.Printf("Average grade with enhanced for loop is: %.2f\n", avg)

	// declaring an array with values in shorthand
	shorthandGrades := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(shorthandGrades)
}
