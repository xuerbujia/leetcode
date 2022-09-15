package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	s := strings.Split(sentence, " ")
	for i, v := range s {
		if len(v) >= len(searchWord) {
			if v[:len(searchWord)] == searchWord {
				return i + 1
			}
		}
	}
	return -1
}
func test1() (t int) {
	t = 1
	defer func() {
		t++
	}()
	return
}
func main() {
	s := test1()
	fmt.Println(s)
}
