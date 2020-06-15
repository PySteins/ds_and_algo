package leetcode

// https://leetcode-cn.com/problems/permutations/
// 46. 全排列

func Permute(nums []int) [][]int {
	var (
		res [][]int
	)
	if len(nums) == 0 {
		return res
	}

	com := make([]int, len(nums))
	used := make([]bool, len(nums))
	dfsP(0, nums, com, used, &res)
	return res
}

func dfsP(idx int, nums, com []int, used []bool, res *[][]int) {
	if idx == len(nums) {
		r := make([]int, len(nums))
		for i, v := range com {
			r[i] = v
		}
		*res = append(*res, r)
		return
	}

	for i, num := range nums {
		if used[i] {
			continue
		}
		com[idx] = num
		used[i] = true
		dfsP(idx+1, nums, com, used, res)
		used[i] = false
	}
}
