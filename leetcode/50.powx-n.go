package leetcode

// https://leetcode-cn.com/problems/powx-n/
// 50. Pow(x, n)
func MyPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == -1 {
		return 1 / x
	}
	odd := false
	if n&1 == 1 {
		odd = true
	}
	half := MyPow2(x, n>>1)

	if odd {
		return half * half * x
	}
	return half * half
}

func MyPow(x float64, n int) float64 {
	neg := false
	if n < 0 {
		neg = true
	}

	y := n
	if neg {
		y = -n
	}

	res := float64(1)
	for y > 0 {
		if y&1 == 1 {
			res *= x
		}
		x *= x
		y = y >> 1
	}
	if neg {
		return 1 / res
	}
	return res
}
