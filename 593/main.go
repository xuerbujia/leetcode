package main

import (
	"fmt"
	"math"
)

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	var p = [][]int{p1, p2, p3, p4}
	var temp []float32
	var khash = map[float32][2]int{}
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			k := getK(p[i], p[j])
			temp = append(temp, k)
			khash[k] = [2]int{i, j}
		}
	}
	var hash = map[float32]int{}
	for _, v := range temp {
		hash[v]++
	}
	//fmt.Println(temp)
	temp = make([]float32, 0, 2)
	var length float64
	for i, v := range hash {
		if v != 2 {
			temp = append(temp, i)
			q := khash[i]
			x := p[q[0]][0] - p[q[1]][0]
			y := p[q[0]][1] - p[q[1]][1]
			l := math.Sqrt(float64(x*x + y*y))
			//	fmt.Println(q, l)
			if length != 0 {
				if length != l {
					return false
				}
			} else {
				length = l
			}
		}
	}
	if len(temp) != 2 {
		return false
	}
	if temp[0] == 0 || temp[1] == 0 {
		if temp[0] == math.MaxFloat32 || temp[1] == math.MaxFloat32 {
			return true
		} else {
			return false
		}
	}
	//getK(p1,p2)
	//getK(p1,p3)
	//getK(p1,p4)
	//getK(p2,p3)
	//getK(p3,p4)

	return int(temp[0]*temp[1]-0.0000001) == -1
}
func getK(x, y []int) float32 {
	if y[0] == x[0] && y[1] == x[1] {
		return math.MinInt64
	}
	if y[0]-x[0] == 0 {
		return math.MaxFloat32
	}
	return float32(y[1]-x[1]) / float32(y[0]-x[0])
}
func main() {
	p1 := []int{1, 1}
	p2 := []int{5, 3}
	p3 := []int{3, 5}
	p4 := []int{7, 7}
	fmt.Println(validSquare(p1, p2, p3, p4))

}
