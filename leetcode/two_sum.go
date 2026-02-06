/**
Two Sum in Go

Problem description: https://leetcode.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	complement := map[int]int{} // map complement to int indice

	for index, val := range nums {
		idx, ok := complement[val]
		if !ok {
			complement[target-val] = index
		} else {
			return []int{index, idx}
		}
	}
	return []int{}
}