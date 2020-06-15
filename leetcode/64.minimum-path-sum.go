package leetcode

// https://leetcode-cn.com/problems/minimum-path-sum/
// 64. 最小路径和

func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	for col := 1; col < cols; col++ {
		grid[0][col] = grid[0][col-1] + grid[0][col]
	}

	for row := 1; row < rows; row++ {
		grid[row][0] = grid[row-1][0] + grid[row][0]
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			grid[row][col] = minInt(grid[row][col-1], grid[row-1][col]) + grid[row][col]
		}
	}
	return grid[rows-1][cols-1]
}

func minPathSum2(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dp := make([]int, cols)
	dp[0] = grid[0][0]

	for col := 1; col < cols; col++ {
		dp[col] = dp[col-1] + grid[0][col]
	}

	left := dp[0]
	for row := 1; row < rows; row++ {
		left += grid[row][0]
		dp[0] = left
		for col := 1; col < cols; col++ {
			dp[col] = minInt(dp[col-1], dp[col]) + grid[row][col]
		}
	}
	return dp[cols-1]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
