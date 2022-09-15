package main

import "math"

func maxProfit(k int, prices []int) int {
	var dp = make([][][2]int, len(prices)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, k+1)
	}
	var minInt = math.MinInt
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i][1] = minInt
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0][1] = minInt
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			//1 当前未持有 说明昨天未持有今天不操作 或者 昨天持有今天卖出
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			//2 当前持有 说明昨天持有今天不操作 或者 昨天未持有今天买入
			dp[i][j][i] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}
	return dp[len(dp)][k][0]
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func main() {
	maxProfit(2, []int{3, 2, 6, 5, 0, 3})
}
