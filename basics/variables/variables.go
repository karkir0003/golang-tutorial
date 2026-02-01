package main

import "fmt"

func main() {
	// one way to declare variable
	var x string
	x = "Hello, World"
	fmt.Println(x)

	// another way to do the above
	var y string = "Hello, World"
	fmt.Println(y)

	// simple concatenation
	var first = "First"
	first = first + " Second"
	fmt.Println(first)

	// Shorthand (COMMON way of creating variable)
	// type is autoinferred by the compiler
	shorthandHelloWorld := "Hello, World"
	fmt.Println(shorthandHelloWorld)

	// Equality of strings
	fmt.Println(x == shorthandHelloWorld) // equality determined by value comparison

	// Constants
	const helloWorldConstant = "Hello, World"
	fmt.Println(helloWorldConstant)
	// CANNOT reassign and you will get compiler error const helloWorldConstant = "Some other string"

	// Multiple variable shorthand. Type is still auto inferred
	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)

}
