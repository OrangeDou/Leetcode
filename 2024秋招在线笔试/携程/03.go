package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	result := make([]int, 0)
	generate(n, m, []int{}, &result)
	count := help(result, k)
	fmt.Print(count)

}

func generate(n, m int, cur []int, result *[]int) {
	if len(cur) == m {
		num := 0
		for i := 0; i < m; i++ {
			num = num*10 + cur[i]
		}
		*result = append(*result, num)
		return
	}
	for i := 0; i <= n; i++ {
		if !contains(cur, i) {
			generate(n, m, append(cur, i), result)
		}
	}
}

func help(nums []int, k int) int {
	count := 0
	for _, num := range nums {
		if num > k {
			count++
		}
	}
	return count

}
func contains(s []int, ele int) bool {
	for _, v := range s {
		if v == ele {
			return true
		}

	}
	return false
}
