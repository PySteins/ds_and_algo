package leetcode

import (
	"strings"
)

// https://leetcode-cn.com/problems/reverse-words-in-a-string/
// 151. 翻转字符串里的单词

func reverseWords(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	// 转为字符数组
	chars := make([]string, l)
	for i, sp := range s {
		chars[i] = string(sp)
	}

	cur := 0
	space := true
	for i, char := range chars {
		if char != " " {
			chars[cur] = chars[i]
			space = false
			cur++
		} else if space == false {
			chars[cur] = " "
			cur++
			space = true
		}
	}
	var validLen int
	if space {
		validLen = cur - 1
	} else {
		validLen = cur
	}
	if validLen < 0 {
		return ""
	}

	reverseWord(chars, 0, validLen)

	prevSpaceIdx := -1
	for i := 0; i < validLen; i++ {
		if chars[i] != " " {
			continue
		}
		reverseWord(chars, prevSpaceIdx+1, i)
		prevSpaceIdx = i
	}
	reverseWord(chars, prevSpaceIdx+1, validLen)

	return strings.Join(chars[:validLen], "")
}

func reverseWord(chars []string, li, ri int) {
	ri--
	for li < ri {
		tmp := chars[li]
		chars[li] = chars[ri]
		chars[ri] = tmp
		li++
		ri--
	}
}
