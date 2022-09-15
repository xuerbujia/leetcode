package main

import (
	"fmt"
)

func movingCount(m int, n int, k int) int {
	var ans int
	var hash = make([][]bool, m)
	for i := range hash {
		hash[i] = make([]bool, n)
	}
	hash[0][0] = true
	var dx = []int{-1, 0}
	var dy = []int{0, -1}
	for i := 0; i < m; i++ {
		temp := i
		var row int
		for temp != 0 {

			row += temp % 10
			temp /= 10
		}

		for j := 0; j < n; j++ {
			temp = j
			var col int
			for temp != 0 {

				col += temp % 10
				temp /= 10
			}

			if row+col <= k {
				//fmt.Println(i, j)
				for c := 0; c < 2; c++ {
					mx := i + dx[c]
					my := j + dy[c]
					if mx < 0 || my < 0 || mx > m-1 || my > n-1 {
						continue
					}
					if hash[mx][my] {
						ans++
						hash[i][j] = true
						break
					}
				}
			}
		}
	}

	return ans + 1
}
func main() {
	fmt.Println(movingCount(38, 15, 9))
}
