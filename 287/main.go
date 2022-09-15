package main

import "fmt"

func findDuplicate(nums []int) int {
	var ans int
	for _, v := range nums {
		ans ^= v
	}
	fmt.Println(ans)
	return 0
}
func main() {
	findDuplicate([]int{1, 3, 4, 2, 2})
}
