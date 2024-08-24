package main

import (
	"fmt"
)

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	sum := make([][]int, 0)
	for i := 0; i <= k; i++ {
		for j := 0; j <= k; j++ {
			for m := k; m >= 0; m-- {
				if i+j+m == k {
					res := make([]int, 0)
					res = append(res, i)
					res = append(res, j)
					res = append(res, m)
					sum = append(sum, res)
				} else {
					continue
				}
			}
		}
	}
	var max int = 0
	for l := 0; l < len(sum); l++ {
		curRes := (a + sum[l][0]) * (b + sum[l][1]) * (c + sum[l][2])
		if curRes > max {
			max = curRes
		}
	}
	fmt.Print(max)
}
