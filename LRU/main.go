package main

import "fmt"

type LRUCache struct {
	Cache    map[int]*Node
	Head     *Node
	Tail     *Node
	MaxCount int
}
type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cache:    make(map[int]*Node),
		MaxCount: capacity,
	}
}
func (c *LRUCache) MoveToHead(node *Node) {
	c.MovePrevAndNextPointer(node)
	c.InsertToHead(node)

}
func (c *LRUCache) MovePrevAndNextPointer(node *Node) {
	next := node.Next
	prev := node.Prev
	if prev == nil {
		fmt.Println(node)
	}
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
	if node == c.Tail {
		c.Tail = prev
	}
}
func (c *LRUCache) InsertToHead(node *Node) {
	if c.Head == nil {
		c.Head = node
		c.Tail = node
		return
	}
	node.Next = c.Head
	node.Prev = nil
	c.Head.Prev = node
	c.Head = node
}
func (c *LRUCache) Delete() {
	tail := c.Tail
	c.MovePrevAndNextPointer(tail)

	delete(c.Cache, tail.Key)
}
func (c *LRUCache) Get(key int) int {
	node := c.Cache[key]
	if node == nil {
		return -1
	}
	if node != c.Head {
		c.MoveToHead(node)
	}
	return node.Val
}
func (c *LRUCache) Put(key int, value int) {
	node := c.Cache[key]
	if node != nil {
		node.Val = value
		if node != c.Head {
			c.MoveToHead(node)
		}
		return
	} else {
		node = &Node{Key: key, Val: value}
		c.Cache[key] = node
		c.InsertToHead(node)
		if len(c.Cache) > c.MaxCount {
			c.Delete()
		}
	}
}
func main() {
	c := Constructor(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	c.Get(4)
	c.Get(3)
	c.Get(2)
	c.Get(1)
	c.Put(5, 5)
}
