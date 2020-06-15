package leetcode

// https://leetcode-cn.com/problems/edit-distance/
// 72. 编辑距离

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	// 转为字符数组
	chars1 := make([]string, l1)
	for i, sp := range word1 {
		chars1[i] = string(sp)
	}
	chars2 := make([]string, l2)
	for i, sp := range word2 {
		chars2[i] = string(sp)
	}

	dp := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for row := 1; row < l1+1; row++ {
		dp[row][0] = row
	}
	for col := 1; col < l2+1; col++ {
		dp[0][col] = col
	}

	for row := 1; row < l1+1; row++ {
		for col := 1; col < l2+1; col++ {
			temp := minD(dp[row][col-1]+1, dp[row-1][col]+1)
			dp[row][col] = minD(temp, dp[row-1][col-1]+1)
			if word1[row-1] == word2[col-1] {
				dp[row][col] = minD(dp[row-1][col-1], dp[row][col])
			}
		}
	}
	return dp[l1][l2]
}

func minD(x, y int) int {
	if x < y {
		return x
	}
	return y
}
