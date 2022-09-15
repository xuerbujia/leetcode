package main

import "fmt"

type Node struct {
	Val  string
	Next *Node
	Prev *Node
}
type BrowserHistory struct {
	head *Node
	now  *Node
}

func Constructor(homepage string) BrowserHistory {
	node := &Node{homepage, nil, nil}

	return BrowserHistory{node, node}
}

func (h *BrowserHistory) Visit(url string) {
	node := &Node{url, nil, nil}
	prev := h.now
	prev.Next = node
	node.Prev = prev
	node.Next = nil
	h.now = node
}

func (h *BrowserHistory) Back(steps int) string {
	node := h.now
	for i := 0; i < steps && node.Prev != nil; i++ {
		node = node.Prev

	}
	h.now = node
	return node.Val
}

func (h *BrowserHistory) Forward(steps int) string {
	node := h.now
	for i := 0; i < steps && node.Next != nil; i++ {
		node = node.Next
	}
	h.now = node
	return node.Val
}
func main() {
	h := Constructor("leetcode")
	h.Visit("google")
	h.Visit("facebook")
	fmt.Println(h.Back(1))
	fmt.Println(h.Back(1))
	h.Visit("linkedin")
	fmt.Println(h.Forward(2))
	h.Back(2)
	h.Back(7)
}
