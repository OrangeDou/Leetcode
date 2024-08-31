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
