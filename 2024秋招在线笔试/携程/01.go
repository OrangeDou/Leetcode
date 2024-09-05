package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	result := help(n, k)
	fmt.Print(result)
}

func help(n, k int) []int {
	p := make([]int, 0)
	for i := 0; i < n; i++ {
		p = append(p, i+1)
	}
	if k == 1 {
		return p
	}

	for i := 0; i < n-k+1; i++ {
		p[i] = i + 1
	}
	for i := n - k + 1; i < n; i++ {
		p[i] = n - (i - (n - k + 1))
	}
	return p
}
