package main

/**
144. 二叉树的前序遍历
https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

*/

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	preorderRecursive(root, &result)
	return result
}

func preorderRecursive(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		preorderRecursive(root.Left, result)
		preorderRecursive(root.Right, result)
	}
}
