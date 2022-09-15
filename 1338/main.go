package main

import "container/heap"

type hp []int

func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Less(i, j int) bool {
	hp := *h
	return hp[i] > hp[j]
}
func (h *hp) Swap(i, j int) {
	hp := *h
	hp[i], hp[j] = hp[j], hp[i]
}
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *hp) Pop() interface{} {
	old := *h
	res := old[old.Len()-1]
	*h = old[:old.Len()-1]
	return res
}
func minSetSize(arr []int) int {
	var count = map[int]int{}
	for _, v := range arr {
		count[v]++
	}
	h := &hp{}
	heap.Init(h)
	for _, v := range count {
		heap.Push(h, v)
	}
	var l = len(arr)
	var ans int
	for h.Len() > 0 {
		i := heap.Pop(h).(int)
		l -= i
		ans++
		if l < len(arr)/2 {
			break
		}
	}
	return ans
}
func main() {

}
