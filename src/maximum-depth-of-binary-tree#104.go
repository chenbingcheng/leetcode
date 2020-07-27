package main

import "container/list"

/**
104. 二叉树的最大深度
https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

*/

func maxDepthByBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushFront(root)

	result := 0

	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		result++
	}
	return result
}

func maxDepthByDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthByDfs(root.Left), maxDepthByDfs(root.Right)) + 1

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
