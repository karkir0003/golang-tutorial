/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Problem: https://leetcode.com/problems/binary-tree-right-side-view/

func rightSideView(root *TreeNode) []int {
	/**
	  Perform level order traversal but after completing a level, you
	  get the last entry in the array
	*/
	result := []int{}
	if root == nil {
		return result
	}
	currLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}
	var lastElement *TreeNode

	for len(currLevel) > 0 {
		// get the last element and populate in results
		lastElement = currLevel[len(currLevel)-1]
		result = append(result, lastElement.Val)

		// go through each element and populate the neighbors

		/**
		  you have to create a fresh array here because when you do   slicing on slices, it's just pointer manipulation on the same underlying array
		*/
		nextLevel = []*TreeNode{}

		for _, node := range currLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		// set curr level to next level
		currLevel = nextLevel
	}
	return result

}