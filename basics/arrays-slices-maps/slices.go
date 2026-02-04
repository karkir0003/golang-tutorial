package main

import "fmt"

func main() {
	// initializing a slice
	initial_slice := make([]float64, 5, 10)
	initial_slice[2] = 20.06
	fmt.Println(initial_slice)

	// taking subslice (start at starting index, end up to but not including end)
	subslice := initial_slice[2:4]
	fmt.Println(subslice)
	start_only_subslice := initial_slice[1:]
	end_only_subslice := initial_slice[:3]
	fmt.Println(start_only_subslice)
	fmt.Println(end_only_subslice)

	// append element to slice
	slice1 := []int{1, 2, 3}       // array of ints with capacity 3!
	slice2 := append(slice1, 4, 5) // new reference because we need to copy over and make new array with different pointer
	fmt.Println("slice 1: ", slice1)
	fmt.Println("slice 2: ", slice2)

	// what happens if we modify entry in slice 1. Are slice 1 and slice 2 referentially different?
	fmt.Println("slice 1 and slice 2 point to same address: ", &slice1[0] == &slice2[0])

	// What happens if we modify the first element of both slices
	slice1[0] = 1900
	fmt.Println("slice 1's first element: ", slice1[0])
	fmt.Println("slice 2's second element: ", slice2[0])

	// append element to slice, but with enough capacity
	slice1 = make([]int, 2, 10) // make an "arraylist" of ints with current length 2, capacity 10
	slice1[0] = 1900
	slice2 = append(slice1, 4, 5)
	fmt.Println("slice 1 originally: ", slice1)
	fmt.Println("slice 2 originally: ", slice2)

	slice1[0] = 500 // set the first element to 500. What happens to slice1 and slice2

	// NOTE: We use &slice1[0] to represent address of first element in the underlying array for the slice
	// &slice1 means the "memory address in the STACK" where the slice1 data lives
	fmt.Println("slice 1 and slice 2 point to same address: ", &slice1[0] == &slice2[0])
	fmt.Println("slice 1: ", slice1)
	fmt.Println("slice 2: ", slice2) // same underlying array for slice2 built on top b/c enough capacity

}
