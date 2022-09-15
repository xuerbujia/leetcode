package main

import "math"

func maxProfit(prices []int, fee int) int {
	var dp = make([][2]int, len(prices)+1)
	dp[0][1] = math.MinInt
	//有手续费 相当于 买入的价格提高了
	for i := 1; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1]-fee)
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
