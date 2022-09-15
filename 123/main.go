package main

import "math"

func maxProfit(prices []int) int {
	var dp = make([][3][2]int, len(prices)+1)
	minInt := math.MinInt
	for i := 0; i < len(dp); i++ {
		dp[i][0][1] = minInt
	}
	for i := 0; i < 3; i++ {
		dp[0][i][1] = minInt
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < 3; j++ {
			//上一天未持有 今天不操作  上一天持有今天卖出
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			//上一天 持有 今天不操作  上一天未持有 今天买入
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}
	return dp[len(dp)-1][2][0]
}
func max(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
func main() {

}
