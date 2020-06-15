package leetcode

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
// 448. 找到所有数组中消失的数字

func findDisappearedNumbers(nums []int) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] *= -1
		}
	}
	var res []int
	for i, num := range nums {
		if num > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
