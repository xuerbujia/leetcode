package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var l int
	for node := head; node != nil; node = node.Next {
		l++
	}
	m := l % k
	l = l / k

	if l == 0 {
		l = 1
		m = 0
	}
	node := head
	var ans = make([]*ListNode, k)
	for i := 0; i < k; i++ {
		ans[i] = node
		extra := 0
		if m > 0 {
			extra = 1
		}
		for j := 0; j < l-1+extra; j++ {
			node = node.Next
		}
		if node == nil {
			break
		}
		next := node.Next
		node.Next = nil
		node = next
		m--
	}

	return ans
}
func main() {
	var s = new(ListNode)
	for i, node := 0, s; i < 10; i++ {
		node.Next = &ListNode{i + 1, nil}
		node = node.Next
	}
	b := splitListToParts(s.Next, 3)
	fmt.Println(b)
}
