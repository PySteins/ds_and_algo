package leetcode

// https://leetcode-cn.com/problems/move-zeroes/
// 283. 移动零

func moveZeroes(nums []int) {
	var cur = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cur], nums[i] = nums[i], nums[cur]
			cur++
		}
	}
}
