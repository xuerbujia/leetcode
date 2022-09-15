package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	var queue []*TreeNode
	var index []int64
	queue = append(queue, root)
	index = append(index, 1)
	var ans int64
	for len(queue) > 0 {
		size := len(queue)
		var min int64 = math.MaxInt64
		var max int64
		for i := 0; i < size; i++ {
			r := queue[0]
			queue = queue[1:]
			idx := index[0]
			index = index[1:]
			if idx < min {
				min = idx
			}
			if idx > max {
				max = idx
			}
			if r.Left != nil {
				queue = append(queue, r.Left)
				index = append(index, idx*2)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
				index = append(index, idx*2+1)
			}
		}
		if ans < max-min+1 {
			ans = max - min + 1
		}
	}
	return int(ans)
}
func main() {
	root := &TreeNode{0, nil, nil}
	root.Left = &TreeNode{0, nil, nil}
	root.Right = &TreeNode{0, nil, nil}
	for i, node := 0, root.Left; i < 25; i++ {
		node.Right = &TreeNode{0, nil, nil}
		node = node.Right
	}
	for i, node := 0, root.Right; i < 25; i++ {
		node.Left = &TreeNode{0, nil, nil}
		node = node.Left
	}
	fmt.Println(math.MaxInt32)
	fmt.Println(widthOfBinaryTree(root))
}
