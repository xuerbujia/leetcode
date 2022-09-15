package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var queue = []*TreeNode{root}
	var findFather = map[*TreeNode]*TreeNode{}
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		if r.Left != nil {
			queue = append(queue, r.Left)
			findFather[r.Left] = r
		}
		if r.Right != nil {
			queue = append(queue, r.Right)
			findFather[r.Right] = r
		}
	}
	var hash = map[*TreeNode]bool{}
	queue = append(queue, target)
	hash[target] = true
	for i := 0; i < k; i++ {
		size := len(queue)
		for s := 0; s < size; s++ {
			r := queue[0]
			queue = queue[1:]
			for j := 0; j < 3; j++ {
				var next *TreeNode
				switch j {
				case 0:
					next = findFather[r]
				case 1:
					next = r.Left
				case 2:
					next = r.Right
				}
				if next == nil || hash[next] {
					continue
				}
				queue = append(queue, next)
				hash[next] = true
			}
		}

	}
	var ans = make([]int, len(queue))
	for i := 0; i < len(ans); i++ {
		ans[i] = queue[i].Val
	}
	return ans
}
func main() {

}
