package main

import "fmt"

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	var stack []byte
	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && num[i] < stack[len(stack)-1] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, num[i])
		if len(stack) == 1 && stack[len(stack)-1] == '0' {
			stack = stack[:len(stack)-1]
		}
	}

	var start = len(stack) - 1
	if len(stack) != 1 {
		for i := 0; i < len(stack); i++ {
			if stack[i] != '0' {
				start = i
				break
			}
		}
	}
	fmt.Println(stack, start)
	return ""
	//  fmt.Println(stack,start)
	//fmt.Println()
	//return string(stack[start:len(stack)-k])
}
func main() {
	removeKdigits("10001", 4)
}
