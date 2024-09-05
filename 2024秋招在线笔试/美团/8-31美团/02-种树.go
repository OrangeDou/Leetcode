/*

长度无限长的公路上，小妹雇佣了n位工人来种树，每个点最多种一棵树。从左向右数，工人所站的位置是a1,a2,a3,....an。已知每位工人都会将自己所在位置的右侧一段长度的区间种满树，且每位工人种树的长度区间相同。现在小妹希望公路上至少有K棵树，为了节约成本，他希望每位工人种树的区间长度尽可能短。请你算出，工人们的种树区间至少多长，才能使得公路上被种上至少K棵树？
*/

package main

import (
	"fmt"
)

func main() {
	n := 0
	k := 0
	fmt.Scan(&n, &k)
	locate := make([]int, n)
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scan(&num)
		locate[i] = num
	}
	fmt.Print(find(locate, k))
}

func help(a []int, l int) int {
	left, right := 1, l
	for left < right {
		mid := (left + right) / 2
		totalTree := 0
		planted := make(map[int]bool)
		for _, v := range a {
			for i := 0; i < mid; i++ {
				pos := v + 1
				if !planted[pos] {
					planted[pos] = true
					totalTree++
				}
			}
		}
		if totalTree >= l {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func find(a []int, k int) int {
	left, right := 0, len(a)-1
	for left < right {
		mid := (left + right) / 2
		if help(a, mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
