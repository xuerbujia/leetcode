package main

import (
	"fmt"
)

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	var pushidx, popidx int
	for pushidx < len(pushed) {
		for len(stack) > 0 && popped[popidx] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			popidx++
		}
		stack = append(stack, pushed[pushidx])
		pushidx++
	}
	fmt.Println(stack)
	fmt.Println(popped[popidx:])
	for i, v := range popped[popidx:] {

		if v != stack[len(stack)-1-i] {
			return false
		}
	}
	return true
}

type test struct {
	Val int `json:"val"`
}

func main() {

}
