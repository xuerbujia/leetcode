package main

/*
简单模拟
因为pieces[i]中的元素是不可以被移动的，
因此我们对于我们可以定义一个map来找出pieces[i][0]在arr中所对应的idx，
再判断从idx开始的len(pieces[i])个元素是否一一对应，
如果不能一一对应就说明是不可行的直接返回false
当全部元素遍历完以后仍然没有返回false即说明满足题意返回true
时间复杂度O(n),空间复杂度O(n)
*/
func canFormArray(arr []int, pieces [][]int) bool {
	var m = map[int]int{}
	for i, v := range arr {
		m[v] = i
	}
	for _, v := range pieces {
		for i := 0; i < len(v); i++ {
			if m[v[0]]+i >= len(arr) || arr[m[v[0]]+i] != v[i] {
				return false
			}
		}
	}
	return true
}
func main() {

}
