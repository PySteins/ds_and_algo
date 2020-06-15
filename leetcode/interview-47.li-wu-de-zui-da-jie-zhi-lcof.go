package leetcode

// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
// 面试题47. 礼物的最大价值

func maxValue(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	dp[0][0] = grid[0][0]
	for col := 1; col < cols; col++ {
		dp[0][col] = dp[0][col-1] + grid[0][col]
	}

	for row := 1; row < rows; row++ {
		dp[row][0] = dp[row-1][0] + grid[row][0]
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			dp[row][col] = maxInt(dp[row-1][col], dp[row][col-1]) + grid[row][col]
		}
	}
	return dp[rows-1][cols-1]
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
