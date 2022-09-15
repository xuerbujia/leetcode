package main

func groupThePeople(groupSizes []int) [][]int {
	var hash = map[int][]int{}
	for i, v := range groupSizes {
		hash[v] = append(hash[v], i)
	}
	var ans [][]int
	for i, v := range hash {
		for j := 0; j < len(v)/i; j++ {
			elem := make([]int, i)
			for k := 0; k < i; k++ {
				elem[k] = v[j*i+k]
			}
			ans = append(ans, elem)
		}
	}
	//fmt.Println(hash)
	//fmt.Println(ans)
	return ans
}
func main() {
	groupThePeople([]int{3, 3, 3, 3, 3, 1, 3})
}
