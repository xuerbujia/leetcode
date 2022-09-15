package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	s := strconv.Itoa(num)
	var dp = make([][2]int, len(s)+1)
	dp[0][0] = 1
	for i := 0; i < len(s); i++ {
		dp[i+1][0] = dp[i][0] + dp[i][1]
		if i >= 1 && s[i] != 0 {
			temp := s[i] - '0' + (s[i-1]-'0')*10
			//	fmt.Println(temp)
			if temp <= 25 {
				dp[i+1][1] = dp[i][0]
			}
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1][0] + dp[len(dp)-1][1]
}
func main() {
	translateNum(2596996162)
	//fmt.Println(rand.Uint32())
	//fmt.Println(rand.Uint32())
	//fmt.Println(rand.Uint32())
	//fmt.Println(rand.Uint32())
	//fmt.Println(rand.Uint32())
	//fmt.Println(rand.Uint32())
}
