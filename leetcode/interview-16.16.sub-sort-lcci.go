package leetcode

// https://leetcode-cn.com/problems/sub-sort-lcci/
// 面试题 16.16. 部分排序
func subSort(array []int) []int {
	l := len(array)
	start, end := -1, -1
	if l == 0 {
		return []int{start, end}
	}

	max := array[0]
	for i := 1; i < l; i++ {
		if max <= array[i] {
			max = array[i]
		} else {
			end = i
		}
	}

	min := array[l-1]
	for i := l - 2; i >= 0; i-- {
		if min >= array[i] {
			min = array[i]
		} else {
			start = i
		}
	}

	return []int{start, end}
}
