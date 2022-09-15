package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	var hash = map[*TreeNode]bool{}
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if !(root.Val == 1 || left || right) {
			hash[root] = true
		}
		fmt.Printf("%p,%d\n", root, root.Val)
		return root.Val == 1 || left || right
	}

	dfs(root)
	if hash[root] {
		return nil
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		if r.Left != nil {
			if hash[r.Left] {
				r.Left = nil
			} else {
				queue = append(queue, r.Left)
			}
		}
		if r.Right != nil {
			if hash[r.Right] {
				r.Right = nil
			} else {
				queue = append(queue, r.Right)
			}
		}
	}
	return root
}
func main() {
	root := &TreeNode{1, nil, &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}}}
	pruneTree(root)
}
