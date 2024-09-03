package main

import (
	"fmt"
	"sort"
)

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 *
 * @param s string字符串
 * @param k int整型
 * @return string字符串
 */
func mostFrequentSubstring(s string, k int) string {
	n := len(s)
	if n < 1 || n > 100 {
		return ""
	}
	if k <= 0 || k > n {
		return ""
	}
	if k == n {
		return s
	}

	mp := make(map[string]int)
	for i := 0; i <= len(s)-k; i++ {
		curString := string(s[i : i+k])
		mp[curString]++
	}

	maxValue := 0
	var maxResult []string
	for key, value := range mp {
		if value > maxValue {
			maxValue = value
			maxResult = []string{key}
		} else if value == maxValue {
			maxResult = append(maxResult, key)
		}
	}

	sort.Strings(maxResult)
	return maxResult[0]
}
func main() {
	a := "AAaa11a1a0a0a"
	fmt.Print(mostFrequentSubstring(a, 2))
}
