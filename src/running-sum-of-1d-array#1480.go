package main

/**
1480. 一维数组的动态和 https://leetcode-cn.com/problems/running-sum-of-1d-array/

给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
请返回 nums 的动态和。

 */

func runningSum(nums []int) []int {
	length := len(nums)
	for i :=1;i<length;i++ {
		nums[i] = nums[i-1]+ nums[i]
	}
	return nums
}
