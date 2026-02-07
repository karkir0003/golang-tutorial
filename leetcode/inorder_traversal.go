package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inOrderTraversalHelper(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}
	/**
	why this chain of 3 statements?

	in each recursive stack, we solve in order for left, get a result
	then append current node to that result using append
	then we pass that already updated result into the solve for right function
	at the end we have the data for the current subtree, so return it back up
	*/
	result = inOrderTraversalHelper(node.Left, result)
	result = append(result, node.Val)
	result = inOrderTraversalHelper(node.Right, result)
	return result
}
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	result = inOrderTraversalHelper(root, result)
	return result
}
