package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	var helper func(i, j int) bool
	helper = func(i, j int) bool {
		if i >= j {
			return true
		}
		p := i
		for postorder[p] < postorder[j] {
			p++
		}
		m := p
		//fmt.Println(m)
		for postorder[p] > postorder[j] {
			p++
		}

		return p == j && helper(i, m-1) && helper(m, j-1)
	}
	return helper(0, len(postorder)-1)
}

func main() {
	fmt.Println(verifyPostorder([]int{1, 7, 2, 6, 5}))
}
