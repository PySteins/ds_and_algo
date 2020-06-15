package leetcode

import (
	"strings"
)

// https://leetcode-cn.com/problems/generate-parentheses/
// 22. 括号生成

func GenerateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return []string{""}
	}
	chars := make([]string, n<<1)
	dfsG(0, n, n, &res, chars)
	return res
}

func dfsG(idx, leftRest, rightRest int, res *[]string, chars []string) {
	if idx == len(chars) {
		*res = append(*res, strings.Join(chars, ""))
		return
	}

	if leftRest > 0 {
		chars[idx] = "("
		dfsG(idx+1, leftRest-1, rightRest, res, chars)
	}

	if rightRest > 0 && rightRest != leftRest {
		chars[idx] = ")"
		dfsG(idx+1, leftRest, rightRest-1, res, chars)
	}

}
