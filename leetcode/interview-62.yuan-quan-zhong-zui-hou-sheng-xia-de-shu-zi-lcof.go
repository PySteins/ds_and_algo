package leetcode

// https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
// 面试题62. 圆圈中最后剩下的数字

func lastRemaining2(n int, m int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (res + m) % i
	}
	return res
}

func lastRemaining(n int, m int) int {
	if n == 0 {
		return 0
	}
	res := (lastRemaining(n-1, m) + m) % n
	return res
}
