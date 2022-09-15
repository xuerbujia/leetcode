package main

import "fmt"

func generateTheString(n int) string {
	var ans string
	if n%2 == 0 {
		ans += "b"
		for i := 0; i < n-1; i++ {
			ans += "a"
		}
	} else {
		for i := 0; i < n; i++ {
			ans += "a"
		}
	}

	return ans
}
func main() {
	fmt.Println(generateTheString(3))
}
