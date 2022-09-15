package main

import (
	"strconv"
	"strings"
)

type Codec struct {
	preorder []int
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var helper func(root *TreeNode)
	var order []string
	helper = func(root *TreeNode) {
		order = append(order, strconv.Itoa(root.Val))
		helper(root.Left)
		helper(root.Right)
	}

	return strings.Join(order, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	o := strings.Split(data, " ")
	var order = make([]int, len(o))
	for i := 0; i < len(o); i++ {
		val, _ := strconv.Atoi(o[i])
		order[i] = val
	}
	var helper func(order []int, left, right int) *TreeNode
	helper = func(order []int, left, right int) *TreeNode {
		if len(order) < 0 || left > right {
			return nil
		}
		val := order[left]
		var mid = right + 1
		root := &TreeNode{val, nil, nil}
		for i := left + 1; i <= right; i++ {
			if order[i] > val {
				mid = i
				break
			}
		}
		root.Left = helper(order, left+1, mid-1)
		root.Right = helper(order, mid, right)
		return root
	}

	return helper(order, 0, len(order)-1)
}

func main() {
}
