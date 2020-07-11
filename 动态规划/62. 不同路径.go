package leadcode

var result int

var targets [][]int = [][]int{{0, 1}, {1, 0}}

func uniquePaths_1(m int, n int) int {
	result = 0
	unique_helper(0, 0, m, n)
	return result
}

func unique_helper(x_start int, y_start int, m int, n int) {
	if x_start == m-1 && y_start == n-1 {
		result++
		return
	}
	for _, item := range targets {
		new_x_start := x_start + item[0]
		new_y_start := y_start + item[1]
		if new_x_start < m-1 && new_y_start < n-1 {
			unique_helper(new_x_start, new_y_start, m, n)
		}
	}
}

func uniquePaths(m int, n int) int {
	var dp [][]int = make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[j-1][i]
		}
	}
	return dp[m-1][n-1]
}
