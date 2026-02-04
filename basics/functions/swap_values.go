package main

import "fmt"

func swap(x *int, y *int) {
	// standard swap using the temp value just like in other languages
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	/**
	Suppose you have x := 1 and y := 2 and you need to swap the values
	*/
	x := 1
	y := 2

	fmt.Println("BEFORE SWAP")
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	swap(&x, &y)
	fmt.Println("AFTER SWAP")
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

}
