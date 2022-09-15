package main

import "fmt"

func finalPrices(prices []int) []int {
	var stack []int
	var cheap = make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {

		if len(stack) > 0 && prices[i] > stack[len(stack)-1] {
			cheap[i] = stack[len(stack)-1]
			stack = append(stack, prices[i])
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > prices[i] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, prices[i])
			cheap[i] = 0
		}
	}
	fmt.Println(cheap)
	return nil
}
func main() {
	finalPrices([]int{4, 4, 8, 5, 1, 7, 9, 2})
}
