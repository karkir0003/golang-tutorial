package main

import "fmt"

func sum(args ...int) int {
	sum := 0
	for _, number := range args {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, -1, 2, -2, 3, -3))
	fmt.Println(sum())
}
