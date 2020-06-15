package leetcode

// https://leetcode-cn.com/problems/reverse-integer/
// 7. 整数反转

func reverse(x int) int {
	var res int
	for x != 0 {

		res = res*10 + (x % 10)
		if !(-(1<<31) <= res && res <= (1<<31)-1) {
			return 0
		}
		x = x / 10
	}

	return res
}
