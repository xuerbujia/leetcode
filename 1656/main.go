package main

import (
	"encoding/json"
	"fmt"
)

type OrderedStream struct {
	str []string
	ptr int
}
type test struct {
	ID       int    `json:"id"`
	TypeName string `json:"type_name"`
}
type st int

func Constructor(n int) OrderedStream {
	o := OrderedStream{
		str: make([]string, n+1),
		ptr: 1,
	}
	s := `{"id":1012929,"type_name":"集群"}`

	t := json.RawMessage(s)
	fmt.Printf("%T\n", t)
	var a test
	json.Unmarshal(t, &a)

	fmt.Printf("%#v", a)
	return o
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	s.str[idKey] = value
	var ans []string
	for s.ptr < len(s.str) && s.str[s.ptr] != "" {
		ans = append(ans, s.str[s.ptr])
		s.ptr++
	}
	return ans
}

func main() {
	Constructor(5)

}
