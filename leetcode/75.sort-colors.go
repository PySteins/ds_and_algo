package leetcode

// https://leetcode-cn.com/problems/sort-colors/submissions/
// 75. 颜色分类
// 扫描一遍数组通常采用双指针或三指针
func sortColors(nums []int) {
	start, cur, end := 0, 0, len(nums)-1
	for cur <= end {
		switch nums[cur] {
		case 0:
			tmp := nums[cur]
			nums[cur] = nums[start]
			nums[start] = tmp
			start++
			cur++
		case 1:
			cur++
		case 2:
			tmp := nums[cur]
			nums[cur] = nums[end]
			nums[end] = tmp
			end--
		}
	}
}
