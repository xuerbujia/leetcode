package main

import (
	"fmt"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	fmt.Println(pairs)
	now := pairs[0]
	var ans = 1
	for _, v := range pairs[1:] {
		if v[0] > now[1] {
			now = v
			ans++
		}
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

}
