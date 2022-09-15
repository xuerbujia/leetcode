package main

import (
	"sort"
)

func arrayRankTransform(arr []int) []int {
	var nums = make([]int, len(arr))
	copy(nums, arr)
	sort.Ints(nums)
	var hash = map[int]int{}
	var count = 1
	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			hash[v] = count
			count++
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = hash[arr[i]]
	}
	//fmt.Println(arr)
	return arr
}
func main() {
	arrayRankTransform([]int{40, 10, 20, 30})
}
