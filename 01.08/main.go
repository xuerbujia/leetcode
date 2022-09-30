package main

/*
使用两个map记录哪一行和哪一列有0元素出现
重新遍历一遍将需要置零的位置置零即可
*/
func setZeroes(matrix [][]int) {
	var row, col = map[int]bool{}, map[int]bool{}
	for i, v := range matrix {
		for j, v1 := range v {
			if v1 == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for k := range row {
		matrix[k] = make([]int, len(matrix[0]))
	}
	for k := range col {
		for i := 0; i < len(matrix); i++ {
			matrix[i][k] = 0
		}
	}
}
func main() {

}
