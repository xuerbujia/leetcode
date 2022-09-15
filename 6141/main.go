package main

import (
	"fmt"
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var value = map[int]bool{}
	var it1 = map[int]int{}
	var nums []int
	for _, v := range items1 {

		value[v[0]] = true
		nums = append(nums, v[0])

		it1[v[0]] = v[1]
	}
	for _, v := range items2 {
		if !value[v[0]] {
			value[v[0]] = true
			nums = append(nums, v[0])
		}
		it1[v[0]] += v[1]
	}

	sort.Ints(nums)
	//fmt.Println(nums)
	var ans [][]int
	for _, v := range nums {
		ans = append(ans, []int{v, it1[v]})
	}
	return ans
}
func countBadPairs(nums []int) int64 {
	var ans int64
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if j-i != nums[j]-nums[i] {
				ans++
			}
		}
	}
	return ans
}
func main() {
	//items1 := [][]int{{1, 3}, {2, 2}}
	//items2 := [][]int{{7, 1}, {2, 2}, {1, 4}}
	//fmt.Println(mergeSimilarItems(items1, items2))
	fmt.Println(countBadPairs([]int{1, 2, 3, 4, 5}))
}
