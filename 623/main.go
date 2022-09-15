package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	var queue []*TreeNode
	queue = append(queue, root)
	//var findFather = map[*TreeNode]*TreeNode{}
	var deep = 2

	for len(queue) > 0 {
		size := len(queue)

		if deep == depth {
			//fmt.Printf("%#v", queue)
			//	l := len(queue)
			for len(queue) > 0 {
				r := queue[0]
				queue = queue[1:]
				left := r.Left
				node1 := &TreeNode{val, nil, nil}
				r.Left = node1
				node1.Left = left
				right := r.Right
				node2 := &TreeNode{val, nil, nil}
				r.Right = node2
				node2.Right = right

			}
			break
		}
		for i := 0; i < size; i++ {
			r := queue[0]
			queue = queue[1:]

			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
		}
		deep++
	}

	return root
}
func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(addOneRow(root, 1, 2))

}
