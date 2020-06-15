package leetcode

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 3. 无重复字符的最长子串

func LengthOfLongestSubstring(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}

	// 转为字符数组
	chars := make([]string, l)
	for i, sp := range s {
		chars[i] = string(sp)
	}

	li := 0
	max := 0
	prevMap := make(map[string]int)
	prevMap[chars[0]] = 0
	for i := 1; i < l; i++ {

		pi, ok := prevMap[chars[i]]
		if ok && pi >= li {
			li = pi + 1
		}
		prevMap[chars[i]] = i
		if max < i-li+1 {
			max = i - li + 1
		}
	}

	return max
}
