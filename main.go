package main

import "fmt"

func main() {
	// leetcode.Constructor(2)
	fmt.Println(throwEgg(2, 100))

}

func throwEgg(n, m int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < m; i++ {
		dp[0][i] = i
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i][j-1] + 1
			if dp[i][j] > m {
				return j
			}
		}
	}
	return -1
}
