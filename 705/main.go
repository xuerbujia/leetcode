package main

type Node struct {
	Val  int
	Next *Node
}
type MyHashSet struct {
	m []*Node
}

func Constructor() MyHashSet {
	return MyHashSet{m: make([]*Node, 65536)}
}
func (s *MyHashSet) Add(key int) {
	node := &Node{Val: key}
	h := key & (65535)
	if s.m[h] == nil {
		s.m[h] = node
	} else {
		if !s.Contains(key) {
			node.Next = s.m[h]
			s.m[h] = node
		}
	}
}
func (s *MyHashSet) Remove(key int) {
	h := key & (65535)
	if s.m[h] != nil {
		node := &Node{Next: s.m[h]}
		for n := node; n.Next != nil; n = node.Next {
			if n.Next.Val == key {
				n.Next = n.Next.Next
				s.m[h] = node.Next
				break
			}

		}
	}
}
func (s *MyHashSet) Contains(key int) bool {
	h := key & (65535)
	if s.m[h] == nil {
		return false
	} else {
		for node := s.m[h]; node != nil; node = node.Next {
			if node.Val == key {
				return true
			}
		}
		return false
	}
}
func main() {
	myHash := Constructor()
	myHash.Add(1)
	myHash.Add(2)
	myHash.Remove(2)
	myHash.Contains(2)
}
