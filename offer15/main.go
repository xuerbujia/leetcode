package main

import "fmt"

func hammingWeight(num uint32) int {
	var ans int
	for num > 0 {
		if num&1 == 1 {
			ans++
		}
		num >>= 1
	}
	return ans
}
func main() {
	fmt.Println(hammingWeight(4294967293))
}
