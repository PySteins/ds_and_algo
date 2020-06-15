package leetcode

import (
	"strings"
)

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// 17. 电话号码的字母组合

var nums = [][]string{
	{"a", "b", "c"}, {"d", "e", "f"},
	{"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"},
	{"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}

	chars := make([]string, len(digits))
	dfs(0, digits, &result, chars)
	return result
}

func dfs(i int, digits string, result *[]string, chars []string) {
	if i == len(digits) {
		// append 会重新分配内存 需要使用指针引用
		*result = append(*result, strings.Join(chars, ""))
		return
	}
	for _, c := range nums[digits[i]-'2'] {
		chars[i] = c
		dfs(i+1, digits, result, chars)
	}
}
