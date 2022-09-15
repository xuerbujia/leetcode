package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var parentsP, parentsQ []*TreeNode
	var dfs func(root *TreeNode, target int, parent *[]*TreeNode)
	dfs = func(root *TreeNode, target int, parent *[]*TreeNode) {
		if root == nil {
			return
		}
		*parent = append(*parent, root)
		if root.Val == target {
			return
		}
		if root.Val > target {
			dfs(root.Left, target, parent)
		} else {
			dfs(root.Right, target, parent)
		}
	}
	dfs(root, p.Val, &parentsP)
	dfs(root, q.Val, &parentsQ)
	var hash = map[*TreeNode]bool{}
	for _, v := range parentsQ {
		hash[v] = true
	}
	for i := len(parentsP) - 1; i >= 0; i-- {
		if hash[parentsP[i]] {
			return parentsP[i]
		}
	}
	return nil
}
func main() {
	p := &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, nil, nil}}
	q := &TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}
	root := &TreeNode{6, p, q}
	fmt.Println(lowestCommonAncestor(root, p, q))
}
