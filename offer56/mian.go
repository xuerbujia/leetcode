package main

import "fmt"

func singleNumbers(nums []int) []int {
	var ans int
	for _, v := range nums {
		ans ^= v
		//	fmt.Println(ans)
	}
	for _, v := range nums {

		fmt.Println(v, v^7)
	}
	fmt.Println(ans, 1^6)
	return nil
}
func main() {
	singleNumbers([]int{4, 1, 4, 6})
}
