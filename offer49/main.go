package main

import (
	"fmt"
	"sync"
)

func nthUglyNumber(n int) int {
	var hash = map[int]bool{}
	var ans int
	var ugly []int
	ugly = append(ugly, 1)
	var i int
	for len(hash) != n {
		for j := 2; j <= 5 && len(hash) != n; j++ {
			if !hash[ugly[i]*j] {
				ugly = append(ugly, ugly[i]*j)
				hash[ugly[i]*j] = true
				ans = ugly[i] * j
			}
			once := sync.Once{}
			once.Do(func() {
				fmt.Println("ff")
			})
		}
		i++
	}
	return ans
}
func main() {
	fmt.Println(nthUglyNumber(10))
}
