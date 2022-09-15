package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var pathQueue [][]int
	var ans int
	//var hash = map[*TreeNode]bool{}
	//pathQueue=append(pathQueue,[]int{root.Val,0})
	var queue []*TreeNode
	queue = append(queue, root)
	pathQueue = append(pathQueue, []int{0})
	for len(queue) > 0 {
		size := len(queue)
		for j := 0; j < size; j++ {
			r := queue[0]
			queue = queue[1:]
			p := pathQueue[0]
			pathQueue = pathQueue[1:]
			t := append([]int{}, p...)
			for i := range t {
				t[i] += r.Val
				if t[i] == targetSum {
					ans++
				}
			}
			t = append(t, 0)

			if r.Left != nil {
				queue = append(queue, r.Left)
				pathQueue = append(pathQueue, t)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
				pathQueue = append(pathQueue, t)
			}
		}
	}
	return ans
}
func main() {

	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	for _, v := range a {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()
			fmt.Println(v)
		}(v)
	}
	wg.Wait()
}
