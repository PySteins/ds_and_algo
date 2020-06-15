package leetcode

// https://leetcode-cn.com/problems/longest-palindromic-substring/
// 5. 最长回文子串

func LongestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}

	var distance, start, max int
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			distance = j - i + 1

			if distance <= 2 && s[i] == s[j] {
				dp[i][j] = true
				if max <= distance {
					max = distance
					start = i
				}
			}

			if distance > 2 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if max < distance {
					max = distance
					start = i
				}
			}
		}
	}
	return s[start : start+max]
}
