package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	memo := make(map[[2]int]int)
	var help func(cur, rem int) int
	help = func(cur, rem int) int {
		if rem < 0 {
			return 0

		}
		if rem == 0 {
			return 1
		}
		if val, ok := memo[[2]int{cur, rem}]; ok {
			return val
		}
		count := 0
		if cur%2 == 0 {
			count += help(cur-1, rem-cur)
			count += help(cur+1, rem-cur)
		} else {
			count += help(cur+1, rem-cur)
			count += help(cur-1, rem-cur)
		}
		memo[[2]int{cur, rem}] = count
		return count
	}
	fmt.Print(help(n, k))
}
