package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	var m = map[string]struct{}{}
	for _, v := range emails {
		index := strings.Index(v, "@")
		var s string
		var dots []int
		last := index
		for i := 0; i < index; i++ {
			if v[i] == '+' {
				last = i
				break
			}
			if v[i] == '.' {
				dots = append(dots, i)
			}
		}
		var pre = 0
		for _, dot := range dots {
			s += v[pre:dot]
			pre = dot + 1
		}
		s += v[pre:last]
		s += v[index:]
		m[s] = struct{}{}
	}
	fmt.Println(m)
	return len(m)
}
func main() {
	numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})
}
