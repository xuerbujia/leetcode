package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var helper func(root *TreeNode)
	//var inOrder []int
	var n int
	var ans int
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		n++
		helper(root.Left)
		if n == k {
			ans = root.Val
			return
		}
		//inOrder = append(inOrder, root.Val)
		helper(root.Right)
	}
	helper(root)
	return ans
}
func main() {

}
