package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root   *TreeNode
	queue  []*TreeNode
	idx    int
	isLeft bool
}

func Constructor(root *TreeNode) CBTInserter {
	var isLeft bool

	var queue []*TreeNode
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		r := q[0]
		q = q[1:]
		if r.Left == nil || r.Right == nil {
			if len(queue) == 0 {
				if r.Left == nil {
					isLeft = true
				}
			}
			queue = append(queue, r)

		}
		if r.Left != nil {
			q = append(q, r.Left)
		}
		if r.Right != nil {
			q = append(q, r.Right)
		}
	}
	return CBTInserter{root: root, queue: queue, idx: 0, isLeft: isLeft}
}

func (c *CBTInserter) Insert(val int) int {
	var ans = c.queue[c.idx].Val
	treeNode := &TreeNode{Val: val}
	c.queue = append(c.queue, treeNode)

	if c.isLeft {
		c.queue[c.idx].Left = treeNode

	} else {
		c.queue[c.idx].Right = treeNode
		c.idx++
	}

	c.isLeft = !c.isLeft
	return ans
}

func (c *CBTInserter) Get_root() *TreeNode {

	return c.root
}
func test(a int) int {
	for i := 0; i < 10; i++ {
	}
	return a
}
func main() {
	var a []int
	a = append(a, 1)

	//fmt.Println(a)
}
