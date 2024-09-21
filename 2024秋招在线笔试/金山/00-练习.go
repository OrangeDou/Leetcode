package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	calcu := 1
	for i := 1; i < n+1; i++ {
		prefix := int(math.Pow(float64(m), float64(i)))
		cow := calculateC(i, n)
		calcu += cow * prefix
	}
	result := calcu % 1000000007
	fmt.Print(result)
}

func calculateC(up, down int) int {
	if up == down {
		return 1
	}
	index := up
	var a, b int = 1, 1
	for i := 0; i < index; i++ {
		a *= down
		b *= up
		down--
		up--
	}
	res := a / b
	return res
}
