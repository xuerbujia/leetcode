package main

import "fmt"

func exist(board [][]byte, word string) bool {

	var hash = make([][]bool, len(board))
	for i := range hash {
		hash[i] = make([]bool, len(board[0]))
	}

	var bfs func(row, col, idx int) bool
	var dx = []int{-1, 0, 0, 1}
	var dy = []int{0, 1, -1, 0}
	bfs = func(row, col, idx int) bool {
		if word[idx] != board[row][col] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}
		for i := 0; i < 4; i++ {
			mx := row + dx[i]
			my := col + dy[i]
			if mx < 0 || mx > len(board)-1 || my < 0 || my > len(board[0])-1 || hash[mx][my] {
				continue
			}
			hash[mx][my] = true
			if bfs(mx, my, idx+1) {
				return true
			}
			hash[mx][my] = false
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			hash[i][j] = true
			if bfs(i, j, 0) {
				return true
			}
			hash[i][j] = false
		}
	}
	return false
}
func main() {
	board := [][]byte{{'a', 'a'}}
	word := "aaa"
	fmt.Println(exist(board, word))
}
