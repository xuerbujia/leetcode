package main

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	var getCoins = make([][][]int, amount+1)
	maxInt := math.MaxInt - 1
	for i := 1; i < len(dp); i++ {
		dp[i] = maxInt
	}
	for i := 1; i < len(dp); i++ {
		var money int
		for _, v := range coins {
			money = i - v
			if money < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[money]+1)
		}
		for _, v := range coins {
			money = i - v
			if money < 0 {
				continue
			}

			if dp[money]+1 == dp[i] {

				if getCoins[money] == nil {
					getCoins[i] = append(getCoins[i], []int{v})

				} else {

					for _, v1 := range getCoins[money] {
						getCoins[i] = append(getCoins[i], append(append([]int{}, v1...), v))
					}

				}

			}
		}
	}

	//for i, v := range getCoins {
	//	fmt.Println(i, v)
	//}
	return 0
}
func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
func main() {

}
