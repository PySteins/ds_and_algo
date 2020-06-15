package leetcode

import "strings"

// https://leetcode-cn.com/problems/string-rotation-lcci/
// 面试题 01.09. 字符串轮转

func isFlippedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return strings.Contains(s1+s1, s2)
}
