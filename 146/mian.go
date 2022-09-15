package main

import "fmt"

type LRUCache struct {
	hash map[int]*ListNode
	head *ListNode
	tail *ListNode
	l    int
}
type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
	prev *ListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{hash: make(map[int]*ListNode, capacity), l: capacity}
}

func (this *LRUCache) Get(key int) int {
	if this.hash[key] == nil {
		return -1
	} else {
		node := this.hash[key]
		res := node.Val
		if node != this.head {

			prev := node.prev
			next := node.Next
			if next == nil {
				this.tail = prev
			} else {
				next.prev = prev
			}
			prev.Next = next

			h := this.head
			node.Next = h
			h.prev = node
			this.head = node
		}

		return res
	}
}

func (this *LRUCache) Put(key int, value int) {
	node := &ListNode{Key: key, Val: value}
	if _, ok := this.hash[key]; !ok {
		this.hash[key] = node
	} else {
		this.hash[key].Val = value
		this.Get(key)
		return
	}
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		next := this.head
		next.prev = node
		node.Next = next
		this.head = node
	}
	if len(this.hash) > this.l {
		tail := this.tail
		prev := tail.prev
		delete(this.hash, tail.Key)
		this.tail = prev
		prev.Next = nil
	}
}
func main() {
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)

	c.Put(4, 1)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	//node := c.tail
	//for node != nil {
	//	fmt.Println(node.Key, node.Val)
	//	node = node.prev
	//}
}
