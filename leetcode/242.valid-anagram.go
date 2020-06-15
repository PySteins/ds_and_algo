package leetcode

// https://leetcode-cn.com/problems/valid-anagram/
// 242. 有效的字母异位词

func IsAnagram(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls == 0 || lt == 0 {
		return true
	}
	if lt != ls {
		return false
	}
	chars := make([]int, 26)
	for i := 0; i < ls; i++ {
		chars[int(s[i])-'a']++
	}

	for i := 0; i < lt; i++ {
		chars[int(t[i])-'a']--
		if chars[int(t[i])-'a'] < 0 {
			return false
		}
	}
	return true
}
