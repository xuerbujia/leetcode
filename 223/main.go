package main

import "fmt"

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	var left1, right1, up1, down1, left2, right2, up2, down2 int
	left1, right1, up1, down1 = getPoint(ax1, ay1, ax2, ay2)
	left2, right2, up2, down2 = getPoint(bx1, by1, bx2, by2)
	left := min(left1, left2)
	right := min(right1, right2)
	up := min(up1, up2)
	down := min(down1, down2)
	extra := getArea(up, down, right, left)
	area1 := getArea(up1, down1, right1, left1)
	area2 := getArea(up2, down2, right2, left2)
	fmt.Println(up, down, right, left)
	return area1 + area2 - extra
}
func getArea(up, down, right, left int) int {
	return (up - down) * (right - left)
}
func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
func getPoint(ax1, ay1, ax2, ay2 int) (left1, right1, up1, down1 int) {
	if ax1 < ax2 {
		left1 = ax1
		right1 = ax2
	} else {
		left1 = ax2
		right1 = ax1
	}
	if ay1 < ay2 {
		down1 = ay1
		up1 = ay2
	} else {
		down1 = ay2
		up1 = ay1
	}
	return
}

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}
