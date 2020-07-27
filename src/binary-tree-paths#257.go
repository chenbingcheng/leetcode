package main

import "strconv"

/**
257. 二叉树的所有路径
https://leetcode-cn.com/problems/binary-tree-paths/

给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

*/

func binaryTreePaths(root *TreeNode) []string {

	var res []string
	if root == nil {
		return res
	}
	dfs(root, &res, "")
	return res

}

func dfs(root *TreeNode, res *[]string, s string) {
	if root.Left == nil && root.Right == nil {
		s += strconv.Itoa(root.Val)
		*res = append(*res, s)
		return
	}

	s += strconv.Itoa(root.Val)
	s += "->"

	if root.Left != nil {
		dfs(root.Left, res, s)
	}

	if root.Right != nil {
		dfs(root.Right, res, s)
	}
}
