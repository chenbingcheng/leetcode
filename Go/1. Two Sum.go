/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

 

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

 */
package Go

/**
暴力
 */
func twoSum(nums []int, target int) []int {
	for i,v := range nums {
		for k := i+1; k < len(nums);k++ {
			if nums[k] + v == target {
				return []int{i,k}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i,v := range nums {
		if k,ok := m[target -v]; ok {
			return []int{k,i}
		}
		m[v] = i
	}
	return []int{}
}