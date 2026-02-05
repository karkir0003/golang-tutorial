package main

import (
	"container/list"
	"fmt"
)

// function to find number in DLL and return the element node!
// Return first occurrence :)
func findNumber(target int, dll *list.List) *list.Element {
	curr := dll.Front()
	for curr != nil {
		if curr.Value == target {
			return curr
		}
		curr = curr.Next()
	}
	return nil
}

// traverse and print all the node values from left to right
func traverseListForward(dll *list.List) {
	fmt.Println("********* Starting Forward Traversal *********")
	curr := dll.Front()
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next()
	}
	fmt.Println("********* Ending Forward Traversal *********")
}

// traverse DLL and print all node values from right to left
func traverseListReverse(dll *list.List) {
	fmt.Println("********* Starting Reverse Traversal *********")
	curr := dll.Back()
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Prev()
	}
	fmt.Println("********* Ending Reverse Traversal *********")
}

func main() {
	dll := list.New()

	// add some numbers. Here our DLL is 3 <-> 1 <-> 2
	dll.PushBack(1)
	dll.PushBack(2)
	dll.PushFront(3)

	dll2 := list.New()
	ctr := 1
	for ctr <= 10 {
		dll2.PushBack(ctr)
		ctr += 1
	}
	// DLL 2: 1 <-> 2 <-> 3 .... <-> 10
	dll.PushBackList(dll2) // After our DLL is 3 <-> 1 <-> 2 <-> 1 .. <-> 10

	// find the first occurrence of 5
	element_five := findNumber(5, dll)
	fmt.Println("element 5: ", element_five) // here we see memory address of prev and nex with the value

	// insert the number 11 after element 5
	dll.InsertAfter(11, element_five)

	element_eleven := findNumber(11, dll)
	fmt.Println("element 11 after the insertion: ", element_eleven)
	traverseListForward(dll)
	traverseListReverse(dll)
}
