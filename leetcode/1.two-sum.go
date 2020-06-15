package leetcode

// https://leetcode-cn.com/problems/two-sum/
// 1. 两数之和

func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i+j == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{i, v}
		}
		m[nums[i]] = i
	}
	return []int{}
}
