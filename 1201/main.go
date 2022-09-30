package main

// type hp []int
//
//	func (h *hp) Len() int {
//		return len(*h)
//	}
//
//	func (h *hp) Less(i, j int) bool {
//		p := *h
//		return p[i] < p[j]
//	}
//
//	func (h *hp) Swap(i, j int) {
//		p := *h
//		p[i], p[j] = p[j], p[i]
//	}
//
//	func (h *hp) Push(x interface{}) {
//		p := *h
//		p = append(p, x.(int))
//		*h = p
//	}
//
//	func (h *hp) Pop() interface{} {
//		p := *h
//		res := p[len(p)-1]
//		*h = p[:len(p)-1]
//		return res
//	}
//
//	func nthUglyNumber(n int, a int, b int, c int) int {
//		var vec = []int{a, b, c}
//		sort.Ints(vec)
//		var m = map[int]bool{}
//		var h = new(hp)
//		heap.Init(h)
//		heap.Push(h, 1)
//		var ans int
//		for i := 0; i <= n; {
//			ans = heap.Pop(h).(int)
//			for _, v := range vec {
//				num := v * ans
//				if !m[num] {
//					m[num] = true
//					heap.Push(h, num)
//				}
//			}
//		}
//		fmt.Println(ans)
//		return ans
//	}
func main() {

}
