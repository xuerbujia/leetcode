package main

func numSpecial(mat [][]int) int {
	//var ans int
	var row = make([]int, len(mat))
	var col = make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}
	var ans int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 && row[i] == 1 && col[j] == 1 {
				ans++
			}
		}
	}
	return ans
}
func main() {

}
