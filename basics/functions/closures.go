package main

import "fmt"

func makeEvenGenerator() func() int {
	i := 0 // originally a local variable in function to go
	// closure
	return func() (ret int) {
		ret = i // because we are using it in a closure, the compiler moves i from stack to heap!
		i += 2
		return
	} // here we are returning a function we can invoke to generate even numbers
}

func main() {
	nextEven := makeEvenGenerator() // the i variable is persisted
	fmt.Println(nextEven())         // 0
	fmt.Println(nextEven())         // 2
	fmt.Println(nextEven())         // 4
}
