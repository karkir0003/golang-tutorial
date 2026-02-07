
/*
*
https://leetcode.com/problems/search-insert-position/description/?envType=problem-list-v2&envId=m5ao8cb6
*/
func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	low := 0
	high := len(nums) - 1
	middle := 0
	for low <= high {
		middle = (low + high) / 2
		if target < nums[middle] {
			// go left
			high = middle - 1
		} else if target == nums[middle] {
			return middle
		} else {
			// go right
			low = middle + 1
		}
	}
	low, high = high, low // make sure high is on high side, low is on low side since after for loop, the property is not preserved
	return high
}