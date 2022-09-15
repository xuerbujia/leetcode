package main

import "fmt"

type MyLinkedList struct {
	head *Node
	tail *Node
	l    int
}
type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{l: -1}
}

func (l *MyLinkedList) Get(index int) int {
	if index > l.l+1 {
		return -1
	}
	n := l.head
	for i := 0; i < index; i++ {
		n = n.Next
	}
	if n == nil {
		return -1
	}
	return n.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	l.l++
	//fmt.Println(l.l)
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.Next = l.head
	l.head = node
}

func (l *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	l.l++
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	l.tail.Next = node
	l.tail = node
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.l+1 {

		return
	}
	if index <= 0 {
		l.AddAtHead(val)

		return
	}
	if index == l.l+1 {

		l.AddAtTail(val)
		return
	}
	n := l.head
	node := &Node{Val: val}
	for i := 0; i < index-1; i++ {
		n = n.Next
	}
	node.Next = n.Next
	n.Next = node
	l.l++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	l.l--
	if index >= 0 && index <= l.l+1 {
		if index == 0 {
			l.head = l.head.Next
			return
		}
		n := l.head

		for i := 0; i < index-1; i++ {
			n = n.Next
		}
		n.Next = n.Next.Next
	} else {
		l.l++
	}
}
func (l *MyLinkedList) switchFunc(s string, opts ...int) {
	var n1, n2 int
	n1 = opts[0]
	if len(opts) > 1 {
		n2 = opts[1]
	}
	switch s {
	case "addAtHead":
		fmt.Println("addAtHead", n1)
		l.AddAtHead(n1)
	case "get":

		ans := l.Get(n1)
		fmt.Printf("get idx:%d,ans:%d\n", n1, ans)
	case "addAtTail":
		fmt.Println("addAtTail", n1)
		l.AddAtTail(n1)
	case "addAtIndex":
		fmt.Println("addAtIndex", n1, n2)
		l.AddAtIndex(n1, n2)
	default:
		fmt.Println("delete", n1, n2)

		l.DeleteAtIndex(n1)

	}
}
func main() {

	c := Constructor()
	var f = []string{"addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"}

	var nums = [][]int{{1}, {3}, {1, 2}, {1}, {2}, {1}}
	fmt.Println(c, len(f), len(nums))
	for i := 0; i < len(f); i++ {

		c.switchFunc(f[i], nums[i]...)
		node := c.head
		var count = 0
		for node != nil {
			fmt.Printf("%d,", node.Val)
			count++
			node = node.Next
		}
		fmt.Println(count == c.l+1)
	}
	//c.AddAtIndex(5, 0)
	//c.AddAtIndex(4, 0)
	//c.AddAtIndex(3, 0)
	//c.AddAtIndex(2, 0)
	//c.AddAtIndex(1, 0)
	//c.AddAtIndex(0, 0)
	//c.DeleteAtIndex(5)

	//c.AddAtHead(4)
	//c.Get(1)
	//c.AddAtHead(1)
	//c.AddAtHead(5)
	//c.DeleteAtIndex(3)
	//c.AddAtHead(7)
	//c.Get(3)
	////c.AddAtHead(1)
	////c.DeleteAtIndex(4)

	//for c.head != nil {
	//	fmt.Println(c.head.Val)
	//	c.head = c.head.Next
	//}
}
