package leetcode

// https://leetcode-cn.com/problems/permutations-ii/
// 47. 全排列 II

func PermuteUnique(nums []int) [][]int {
	var (
		res [][]int
	)
	if len(nums) == 0 {
		return res
	}

	com := make([]int, len(nums))
	used := make([]bool, len(nums))
	dfsPU(0, nums, com, used, &res)
	return res
}

func dfsPU(idx int, nums, com []int, used []bool, res *[][]int) {
	if idx == len(nums) {
		r := make([]int, len(nums))
		for i, v := range com {
			r[i] = v
		}
		*res = append(*res, r)
		return
	}

	usedMap := make(map[int]struct{})
	for i, num := range nums {
		if used[i] {
			continue
		}
		if _, ok := usedMap[num]; ok {
			continue
		}
		usedMap[num] = struct{}{}
		com[idx] = num
		used[i] = true
		dfsPU(idx+1, nums, com, used, res)
		used[i] = false
	}
}
