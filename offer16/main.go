package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var ans float64 = 1
	for n > 0 {
		if n&1 == 1 {
			ans *= x
		}
		n >>= 1
		x *= x
	}
	return ans
}
func main() {
	fmt.Println(myPow(-2, -3))
}
