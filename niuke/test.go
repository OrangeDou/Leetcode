package main

import (
	"fmt"
)

// 定义地毯结构
type Carpet struct {
	ID      int
	LeftX   int
	BottomY int
	RightX  int
	TopY    int
}

// 检查点是否被地毯覆盖
func (c Carpet) Covers(x, y int) bool {
	return x >= c.LeftX && x <= c.RightX && y >= c.BottomY && y <= c.TopY
}

// 找到覆盖给定点的最上面那张地毯的编号
func FindTopCarpet(carpets []Carpet, x, y int) int {
	for i := len(carpets) - 1; i >= 0; i-- {
		if carpets[i].Covers(x, y) {
			return carpets[i].ID
		}
	}
	return -1 // 如果没有地毯覆盖该点，返回-1
}

func main() {
	var n int
	fmt.Scan(&n) // 读取地毯数量

	carpets := make([]Carpet, n)
	for i := 0; i < n; i++ {
		var a, b, g, k int
		fmt.Scan(&a, &b, &g, &k) // 读取每张地毯的信息
		carpets[i] = Carpet{
			ID:      i + 1,
			LeftX:   a,
			BottomY: b,
			RightX:  a + g,
			TopY:    b + k,
		}
	}

	var x, y int
	fmt.Scan(&x, &y) // 读取查询点的坐标

	topCarpetID := FindTopCarpet(carpets, x, y)
	fmt.Println(topCarpetID)
}
