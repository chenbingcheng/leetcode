package main

/**
1512. 好数对的数目 https://leetcode-cn.com/problems/number-of-good-pairs/

给你一个整数数组 nums 。
如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。
返回好数对的数目。

 */

/**
暴力，不解释
 */
func numIdenticalPairs1(nums []int) int {
	var result = 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				result++
			}
		}
	}
	return result
}


/**
use map
 */
func numIdenticalPairs2(nums []int) int {
	var result int = 0
	var numsMap = make(map[int]int)
	for _,v := range nums {
		numsMap[v] += 1
	}

	for _,value := range numsMap {
		result += value * (value - 1) / 2
	}
	return result
}


