package main

import "container/list"

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func levelOrderByBfs(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	// queue
	queue := list.New()

	queue.PushFront(root)

	for queue.Len() > 0 {
		var current []int
		listLength := queue.Len()
		for i := 0; i < listLength; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)
			current = append(current, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		result = append(result, current)

	}
	return result
}

// dfs
func levelOrderByDfs(root *TreeNode) [][]int {
	return dfs(root, 0, [][]int{})
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}

	if len(res) == level {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}

	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)

	return res
}
