package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	var min int = math.MaxInt
	var ans int
	for _, v := range prices {
		if v < min {
			min = v
			continue
		}
		ans = max(ans, v-min)
	}
	return ans
}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func main() {
	fmt.Println(maxProfit([]int{7, 6, 5, 4, 3, 2, 1}))
}
