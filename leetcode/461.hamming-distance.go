package leetcode

// https://leetcode-cn.com/problems/hamming-distance/
// 461. 汉明距离

func hammingDistance(x int, y int) int {
	count := 0
	for x != y {
		if x&1 != y&1 {
			count++
		}
		x >>= 1
		y >>= 1
	}
	return count
}
