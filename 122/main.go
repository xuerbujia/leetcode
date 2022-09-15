package main

import "math"

func maxProfit(prices []int) int {
	var dp = make([][2]int, len(prices)+1)
	dp[0][1] = math.MinInt
	for i := 1; i < len(dp); i++ {
		//当前未持有 1.上一天未持有今天不操作 2.上一天持有今天卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		//当前持有 1.上一天持有今天不操作 2.上一天不操作今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}
	return dp[len(dp)-1][0]
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
