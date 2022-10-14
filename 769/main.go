package main

func maxChunksToSorted(arr []int) int {

	var max int
	var ans int
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if max == i {
			ans++
		}
	}
	//	fmt.Println(stack)
	//fmt.Println(ans)
	return ans
}
func main() {
	maxChunksToSorted([]int{1, 0, 2, 3, 4})
}
