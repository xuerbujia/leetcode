package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	var s []string
	for _, v := range nums {
		s = append(s, strconv.Itoa(v))
	}
	xbegin := 0
	ybegin := 0
	var less func(i, j int) bool
	less = func(i, j int) bool {
		x := s[i][xbegin:]
		y := s[j][ybegin:]

		if x == y {
			return s[i] < s[j]
		}
		//var xbit, ybit []byte
		var k int
		for k = 0; k < len(x) && k < len(y); k++ {
			//xbit = append(xbit, x[k]-'0'%10)
			//ybit = append(ybit, y[k]-'0'%10)
			if x[k] < y[k] {
				return true
			}
			if x[k] > y[k] {
				return false
			}
		}

		if k == len(x) {
			ybegin += k
			if less(i, j) {
				ybegin = 0
				return false
			} else {
				ybegin = 0
				return true
			}

		}
		if k == len(y) {
			xbegin += k
			if less(i, j) {
				xbegin = 0
				return true
			}
			xbegin = 0
		}

		return false
	}
	sort.Slice(s, less)
	fmt.Println(s)
	var ans string
	for _, v := range s {
		ans += v
	}
	return ans
}
func main() {
	minNumber([]int{12, 121})
}
