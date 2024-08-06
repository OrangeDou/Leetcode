package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) == len(t) {
		if s == t {
			return s
		}
		return ""
	}
	if len(s) < len(t) {
		return ""
	}
	need := make(map[rune]int)
	window := make(map[rune]int)

	for _, c := range t {
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	start, size := 0, len(s) // 记录最小覆盖子串的起始索引及长度
	length := len(s)

	for right < length {
		c := rune(s[right]) // c 是将移入窗口的字符
		right++
		if _, ok := need[c]; ok { // 进行窗口内数据的一系列更新
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) { // 判断左侧窗口是否要收缩
			if (right - left) < size {
				start = left
				size = right - left
			}
			if (right - left) == size {
				return s
			}	
			d := rune(s[left]) // d 是将移出窗口的字符
			left++
			if _, ok := need[d]; ok { // 进行窗口内数据的一系列更新
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if size == len(s) { // 返回最小覆盖子串
		return ""
	}

	return s[start : start+size]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Print(minWindow(s, t))
}
