package main

import "fmt"

func constructArr(a []int) []int {
	var prefix, postfix = make([]int, len(a)), make([]int, len(a))
	prefix[0] = 1
	postfix[len(postfix)-1] = 1
	for i := 1; i < len(a); i++ {
		prefix[i] = prefix[i-1] * a[i-1]
	}
	for i := len(a) - 2; i >= 0; i-- {
		postfix[i] = a[i+1] * postfix[i+1]
	}
	for i := range prefix {
		prefix[i] *= postfix[i]
	}
	fmt.Println(prefix)
	return prefix
}
func main() {
	constructArr([]int{1, 2, 3, 4, 5})
}
