package main

import "fmt"

/*
看到题意说到接下来的K个数字之和和之前K个数字之和首先想到了前缀和
当K>0时使用前缀和所求值为 index=K+i的前缀和减去i的前缀和，
又因为可以看做一个循环的数组，因此当K+i越界时需要对K+i取模后再找前缀和
而一旦发生越界就会导致少加了从idx=len(code)-1的前缀和，需要重新加上

k<0时同理
*/
func decrypt(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}
	var pre = make([]int, len(code)+1)
	var post = make([]int, len(code)+1)
	for i := 1; i < len(pre); i++ {
		pre[i] = pre[i-1] + code[i-1]
	}
	for i := len(post) - 2; i >= 0; i-- {
		post[i] = post[i+1] + code[i]
	}
	pre = pre[1:]
	post = post[:len(post)-1]
	for i := range code {
		if k > 0 {
			code[i] = pre[(k+i)%(len(code))] - pre[i]
			if i+k >= len(code) {
				code[i] += pre[len(pre)-1]
			}
		} else {
			code[i] = pre[(i+k+len(code))%(len(code))] - pre[i]
			if i+k < 0 {
				code[i] += pre[0]
			}
		}
	}
	fmt.Println(code)
	return nil
}
func main() {

}
