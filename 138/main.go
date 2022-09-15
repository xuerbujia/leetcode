package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var idxOfNode = map[*Node]int{}
	var NodeIdxToRandomIdx = map[int]int{}
	node := head
	var idx int
	for node != nil {
		idxOfNode[node] = idx
		idx++
		node = node.Next
	}

	node = head
	idx = 0
	for node != nil {
		var randomIdx = 0
		if node.Random == nil {
			randomIdx = -1
		} else {
			randomIdx = idxOfNode[node.Random]
		}
		NodeIdxToRandomIdx[idx] = randomIdx
		idx++
		node = node.Next
	}

	var ans = new(Node)
	ansnode := ans
	node = head
	findNode := make(map[int]*Node)
	idx = 0
	for node != nil {
		ansnode.Next = &Node{node.Val, nil, nil}
		node = node.Next
		ansnode = ansnode.Next
		findNode[idx] = ansnode
		idx++
	}
	ans = ans.Next
	fmt.Println(ans)
	ansnode = ans
	idx = 0
	for ansnode != nil {
		if NodeIdxToRandomIdx[idx] != -1 {
			ansnode.Random = findNode[idx]
		}
		ansnode = ansnode.Next
	}
	return ans
}
func main() {

}
