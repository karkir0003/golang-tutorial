package main

import "fmt"

func columnOrderTraversal(arr [][]int) []int {
	/**
	Perform column order traversal for a 2D array
	*/
	result := []int{}

	row_ptr := 0
	col_ptr := 0

	for col_ptr < len(arr[0]) {
		for row_ptr < len(arr) {
			result = append(result, arr[row_ptr][col_ptr])
			row_ptr += 1
		}
		col_ptr += 1
		row_ptr = 0
	}
	return result

}

func rowOrderTraversal(arr [][]int) []int {
	/**
	Perform row order traversal for a 2D array
	*/

	result := []int{}

	row_ptr := 0
	col_ptr := 0

	for row_ptr < len(arr) {
		for col_ptr < len(arr[0]) {
			result = append(result, arr[row_ptr][col_ptr])
			col_ptr += 1
		}
		row_ptr += 1
		col_ptr = 0
	}
	return result
}

func main() {
	/**
	Create 2D array of ints and do row order and column order traversal
	*/
	arr1 := [5][5]int{{}} // 5x5 2D array

	num_cols := len(arr1[0])

	// Count in consecutive order
	for row, entry := range arr1 {
		for col, _ := range entry {
			arr1[row][col] = row*num_cols + col
		}
	}

	fmt.Println("array: ", arr1)

	// we need to convert the fixed array to slice
	slice := make([][]int, len(arr1)) // make the slice with capacity as num rows
	for i := range arr1 {
		slice[i] = arr1[i][:] //point each slice element to the corresponding row
	}

	fmt.Println("Row order traversal: ", rowOrderTraversal(slice))
	fmt.Println("Column order traversal: ", columnOrderTraversal(slice))

}
