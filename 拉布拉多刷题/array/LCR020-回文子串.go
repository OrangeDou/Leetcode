package main

import "fmt"

func countSubstrings(s string) int {
	subString := getSubstring(s)
	result := 0
	for _, v := range subString {
		sig := true
		length := len(v)
		left, right := 0, length-1
		for left < right {
			if v[left] == v[right] {
				left++
				right--
			} else {
				sig = false
				break
			}
		}
		if sig {
			result++
		}
	}
	return result
}

func getSubstring(s string) []string {
	resu := make([]string, 0)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			resu = append(resu, s[i:j])
		}
	}
	return resu
}

func main() {
	s := "abc"
	fmt.Print(countSubstrings(s))
}
