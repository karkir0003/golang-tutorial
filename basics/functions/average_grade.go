package main

import "fmt"

func average(grades []float64) float64 {
	totalPoints := 0.0
	for _, grade := range grades {
		totalPoints += grade
	}

	return totalPoints / float64(len(grades))
}

func main() {
	grades := []float64{
		98,
		93,
		77,
		82,
		83,
	}
	fmt.Println("Average: ", average(grades))
}
