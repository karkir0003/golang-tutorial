package main

import "fmt"

func main() {
	// Integer and float operations
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)
	// NOTE: integer division is strict truncation (like 5 // 2 in python)
	fmt.Println("5 / 2 =", 5/2)
	fmt.Println("5.0 / 2 =", 5.0/2)

	// string operations
	fmt.Println("Length of Hello World is: ", len("Hello World"))

	/**
	NOTE: This function extracts the character at a given index in a string. You'll see here that we don't get the
	character itself. Instead you get the byte representation (ASCII) of the character.

	0 indexing just like Python
	*/
	fmt.Println("5th character of Hello World is: ", "Hello World"[4])

	fmt.Println("5th character of Hello World string in string representation: ", string("Hello World"[4]))

	fmt.Println("Take the 'ello' substring of Hello World: ", string("Hello World"[1:5]))

	fmt.Println("Hello plus World =", "Hello"+"World") // direct string concat b/c Go infers the type of the concatenated elements

	// boolean
	fmt.Println("true AND false =", true && false)
	fmt.Println("!false =", !false)
	fmt.Println("true || false =", true || false)

	fmt.Println("true XOR true =", (true || false) && !(true && false))
}
