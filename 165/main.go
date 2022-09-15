package main

import (
	"fmt"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < len(v1) || i < len(v2); i++ {
		var s1, s2 string
		if i < len(v1) {
			s1 = v1[i]
		}
		if i < len(v2) {
			s2 = v2[i]
		}
		var n1, n2 int
		for j := 0; j < len(s1); j++ {
			n1 *= 10
			n1 += int(s1[j] - '0')
		}
		for j := 0; j < len(s2); j++ {
			n2 *= 10
			n2 += int(s2[j] - '0')
		}
		if n1 > n2 {
			return 1
		}
		if n1 < n2 {
			return -1
		}
	}
	//fmt.Printf("%#v\n%#v", v1, v2)
	return 0
}
func main() {
	fmt.Println(compareVersion("1.001.1", "1.01"))
}
