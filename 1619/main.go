package main

import (
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	start := int(float64(len(arr)) * 0.05)
	var sum int
	for i := start; i < len(arr)-1-start; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr)-2*start)
}
func main() {

}
