package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for len(pushed) > 0 {
		if len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		} else {
			stack = append(stack, pushed[0])
			pushed = pushed[1:]
		}
	}
	//fmt.Println(stack, popped)
	for i := range stack {
		if stack[i] != popped[len(popped)-1-i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
