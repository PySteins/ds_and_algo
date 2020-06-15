package leetcode

import "sort"

// https://leetcode-cn.com/problems/majority-element/
// 169. 多数元素

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
