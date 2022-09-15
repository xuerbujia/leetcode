package main

func maxValue(grid [][]int) int {
	var dp = make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 2; i < len(grid); i++ {
		for j := 2; j < len(grid[0]); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func main() {

}
