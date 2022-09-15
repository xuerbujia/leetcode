package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var p = make([]int, len(nums)+1)
	p[len(p)-1] = 0
	for i := len(p) - 2; i >= 0; i-- {
		p[i] = p[i+1] + nums[i]
	}
	fmt.Println(nums)
	fmt.Println(p)
	var sum int
	var ans []int
	for i := 0; i < len(p)-1; i++ {
		if sum > p[i] {
			break
		}
		sum += nums[i]
		ans = append(ans, nums[i])
	}
	fmt.Println(ans)
	return ans
}
func main() {
	minSubsequence([]int{4, 4, 7, 6, 7})
}
