package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	//fmt.Printf("%#v", words)
	var ans string
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == "" {
			continue
		}
		ans += words[i] + " "
	}

	return ans[:len(ans)-1]
}
func main() {
	fmt.Printf("%#v", reverseWords("  hello world  "))
}
