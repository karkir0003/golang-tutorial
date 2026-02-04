package main

import "fmt"

func fibonacci(n int) int {
	if n < 0 {
		panic("Cannot take fibonacci for negative numbers")
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(-10))
}
