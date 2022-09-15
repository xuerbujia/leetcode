package main

import (
	"fmt"
	"strconv"
	"time"
)

func maximumSwap(num int) int {

	s := []byte(strconv.Itoa(num))
	maxIdx, idx1, idx2 := len(s)-1, len(s)-1, -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else if s[i] < s[maxIdx] {
			idx1, idx2 = i, maxIdx
		}
	}
	if idx2 < 0 {
		return num
	}
	s[idx1], s[idx2] = s[idx2], s[idx1]
	num, _ = strconv.Atoi(string(s))
	//fmt.Println(num)
	return num
}
func main() {
	t := time.Now()
	date := time.Date(t.Year(), t.Month(), t.Day(), 18, 0, 0, 0, &time.Location{})
	fmt.Println(date)
}
