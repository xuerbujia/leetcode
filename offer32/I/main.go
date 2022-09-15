package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var queue []*TreeNode
	var ans []int
	queue = append(queue, root)
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		ans = append(ans, r.Val)
		if r.Left != nil {
			queue = append(queue, r.Left)
		}
		if r.Right != nil {
			queue = append(queue, r.Right)
		}
	}
	return ans
}
func main() {

}
