package main

import (
	"fmt"
	"strconv"
)

func fractionAddition(expression string) string {
	var symbols []byte
	var begin int
	if expression[0] == '+' || expression[0] == '-' {
		symbols = append(symbols, expression[0])
		begin = 1
	} else {
		symbols = append(symbols, '+')
	}
	var nums []int
	var num int
	for i := begin; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '/' {
			nums = append(nums, num)
			if expression[i] != '/' {
				symbols = append(symbols, expression[i])
			}
			num = 0
			continue
		}
		num *= 10
		num += int(expression[i] - '0')
	}
	nums = append(nums, num)
	var denominator = 1
	for i := 1; i < len(nums); i += 2 {
		denominator *= nums[i]
	}
	var divisor int
	for i := 0; i < len(symbols); i++ {
		if symbols[i] == '+' {
			//	fmt.Printf("nums:=%d\n", denominator/nums[i*2+1])
			divisor += nums[i*2] * (denominator / nums[i*2+1])
		} else {
			//	fmt.Printf("nums:=%d\n", denominator/nums[i*2+1])
			divisor -= nums[i*2] * (denominator / nums[i*2+1])
		}
	}
	//fmt.Println(nums, symbols)
	//fmt.Println(divisor, denominator)
	g := gcd(divisor, denominator)
	if divisor < 0 {
		g = gcd(-divisor, denominator)
	}
	return strconv.Itoa(divisor/g) + "/" + strconv.Itoa(denominator/g)
}
func gcd(a, b int) int {
	//fmt.Println(a, b)
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func main() {
	fmt.Printf(fractionAddition("5/3+1/3"))
}
