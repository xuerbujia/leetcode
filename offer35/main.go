package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var ans = new(Node)
	node := ans
	headnode := head
	var addressToIndex = map[*Node]int{}
	var i int
	for headnode != nil {
		addressToIndex[headnode] = i
		i++
		headnode = headnode.Next
	}
	fmt.Println(addressToIndex)
	var hash = map[int]int{}
	headnode = head
	i = 0
	for headnode != nil {
		if headnode.Random != nil {
			hash[i] = addressToIndex[headnode.Random]
		}
		i++
		headnode = headnode.Next
	}
	indexToAddress := make(map[int]*Node)
	i = 0
	for head != nil {
		newNode := &Node{
			Val:    head.Val,
			Next:   nil,
			Random: nil,
		}
		indexToAddress[i] = newNode
		i++
		node.Next = newNode
		head = head.Next
		node = node.Next
	}
	for k, v := range hash {
		indexToAddress[k].Random = indexToAddress[v]
	}
	fmt.Println(hash)
	return ans
}
func main() {

}
