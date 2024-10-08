/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "baa"
	p := "aa"
	fmt.Print(findAnagrams2(s, p))
}

func findAnagrams(s string, p string) []int {
	arrayP := make([]string, 0)
	for k := 0; k < len(p); k++ {
		arrayP = append(arrayP, p[k:k+1])
	}

	result := make([]int, 0)
	window := 3
	for i := 0; i < len(s)-3; i++ {
		current := make([]string, 0)
		for j := i; j < window; j++ {
			current = append(current, s[j:j+1])
		}
		sort.Strings(current)
		if hash(current) == hash(arrayP) {
			result = append(result, i)
		}

	}
	return result

}

func hash(s []string) string {
	sort.Strings(s)
	var sb strings.Builder
	for i, v := range s {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(v)
	}
	return sb.String()

}

func findAnagrams2(s string, p string) []int {
	win := make(map[rune]int, 0)
	target := make(map[rune]int, 0)
	tRune := []rune(p)
	for _, c := range tRune {
		target[c]++
	}
	result := make([]int, 0)

	left, right := 0, 0
	valid := 0
	sRune := []rune(s)
	for right < len(sRune) {
		c := sRune[right]
		right++
		if _, ok := target[c]; ok {
			win[c]++
			if win[c] == target[c] {
				valid++
			}
		}
		for right-left >= len(target) {
			if valid == len(target) {
				result = append(result, left)
			}
			d := sRune[left]
			left++
			if _, ok := target[d]; ok {
				if win[d] == target[d] {
					valid--
				}
				win[d]--
			}
		}

	}

	return result
}
