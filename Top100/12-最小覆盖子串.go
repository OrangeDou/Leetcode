package main

import (
	"fmt"
	"math"
	"reflect"
)

func minWindow(s string, t string) int {
	target := make(map[string]int, 0)
	currentResult := make(map[string]int, 0)
	minResult := 9999999999
	for _, v := range t {
		target[string(v)]++
	}

	for i := 0; i < len(s); i++ {
		j := i
		for {
			if _, exists := target[s[j:j+1]]; exists {
				currentResult[s[j:j+1]]++
			}
			j++
			if j == len(s) {
				break
			}
		}
		if mapsAreEqual(currentResult, target) {
			if j+1 < minResult {
				minResult = j + 1
			}
		}

	}
	if minResult == 9999999999 {
		return 0
	}
	return minResult
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Print(minWindows(s, t))
}

func mapsAreEqual(m1, m2 map[string]int) bool {
	return reflect.DeepEqual(m1, m2)
}

//题解 滑动窗口

func minWindows(s, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]

}
