package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	var ans string
	var num = float64(numerator) / float64(denominator)
	var integer = numerator / denominator
	var decimal = num - float64(integer)
	var hash = map[string]bool{}
	var i int
	var s string
	for float64(i) != decimal {
		decimal *= 10
		i = int(decimal)
		decimal -= float64(i)
		s += strconv.Itoa(i)
		if hash[s] {
			break
		}
		hash[s] = true
	}
	fmt.Println(integer, s)
	return ans
}
func main() {
	fractionToDecimal(4, 333)
}
