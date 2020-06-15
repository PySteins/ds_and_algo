package leetcode

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
// 977. 有序数组的平方
// 类似 88. 合并两个有序数组
func sortedSquares(A []int) []int {
	l := len(A)
	newA := make([]int, l)
	start := 0
	end := l - 1
	for start <= end {
		if A[start]*A[start] > A[end]*A[end] {
			newA[end-start] = A[start] * A[start]
			start++
		} else {
			newA[end-start] = A[end] * A[end]
			end--
		}
	}
	return newA
}
