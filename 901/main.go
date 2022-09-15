package main

type StockSpanner struct {
	stack [][2]int
	//prices []int
	//m      map[int]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		//m: make(map[int]int)
	}
}

func (s *StockSpanner) Next(price int) int {
	var count = 1
	for len(s.stack) > 0 && price > s.stack[len(s.stack)-1][0] {
		count += s.stack[len(s.stack)-1][1]
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, [2]int{price, count})
	return s.stack[len(s.stack)-1][0]
}
func main() {

}
