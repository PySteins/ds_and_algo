package leetcode

// https://leetcode-cn.com/problems/valid-parentheses/
// 20. 有效的括号
import (
	"container/list"
)

func isValid(s string) bool {
	p := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	stack := list.New()
	for _, sp := range s {
		str := string(sp)
		if str == "{" || str == "[" || str == "(" {
			stack.PushBack(str)
		} else {
			v := stack.Back()
			if v == nil {
				return false
			}
			if p[str] != v.Value.(string) {
				return false
			}
			stack.Remove(v)
		}
	}
	if stack.Back() != nil {
		return false
	}
	return true
}
