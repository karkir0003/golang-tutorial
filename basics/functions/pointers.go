package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0 // pass in the memory address of x and update the value of it to 0
}

func updateArrayNoChange(arr [3]int) {
	arr[0] = 100 // go is updating a deepcopy of the passed in array, but it's never returned
}

func updateArrayChange(arr *[3]int) {
	(*arr)[0] = 100 // dereference arr using (*arr) FIRST to get array and then access specific index
}

func updateSlice(slice []int) {
	slice[0] = 100 //operating on a copy of pointer to same underlying array in heap
}

func main() {
	x := 5
	fmt.Println("value of x before update: ", x)
	zero(&x) // pass in address of x
	fmt.Println("value of x after update: ", x)

	// update element of array - no change
	array := [3]int{1, 2, 3}
	fmt.Println("array before update: ", array)
	updateArrayNoChange(array)
	fmt.Println("array after updateArrayOne: ", array) // pass by value, so no underlying data changed!

	// update element of array - working update
	updateArrayChange(&array)
	fmt.Println("array after updateArrayChange: ", array)

	// update slice
	slice := []int{1, 2, 3}
	fmt.Println("slice before update: ", slice)
	updateSlice(slice)
	fmt.Println("Slice after updateSlice: ", slice) // value changes b/c still same underlying array
}
