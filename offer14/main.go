package main

import "fmt"

func cuttingRope(n int) int {
	var dp = make([]int64, n+1)
	dp[1] = 1
	var i, j int64
	for i = 2; i < int64(len(dp)); i++ {
		for j = 1; j <= int64(float32(i)/2+0.5); j++ {

			dp[i] = max(dp[i], max(dp[i-j]*j, j*(i-j))%1000000007)
		}
	}

	return int(dp[len(dp)-1])
}
func max(i, j int64) int64 {
	if i > j {
		return i
	} else {
		return j
	}
}
func main() {
	fmt.Println(cuttingRope(120))
}
