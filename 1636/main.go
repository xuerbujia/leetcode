package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	var count = map[int]int{}
	for _, v := range nums {
		count[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return count[a] < count[b] || count[a] == count[b] && a > b
	})
	return nums
}
func main() {
	fmt.Println(frequencySort([]int{1, 2, 2, 3, 1, 1}))
}
