package main

import "fmt"

type CQueue struct {
	in    []int
	out   []int
	count int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
	this.count++
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		}
		for i := 0; i < this.count; i++ {
			temp := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1]
			this.out = append(this.out, temp)
		}
		this.in = make([]int, 0, len(this.in))
	}
	res := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	this.count--
	return res
}

func main() {
	c := Constructor()
	c.AppendTail(5)
	c.AppendTail(2)
	fmt.Println(c.in)
	fmt.Println(c.out)
	fmt.Println(c.DeleteHead())
	fmt.Println(c.out)
	fmt.Println(c.DeleteHead())
	fmt.Println(c.out)
}
