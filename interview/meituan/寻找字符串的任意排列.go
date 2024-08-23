package main

import "fmt"

func main() {
	a := "aba"
	b := "ngbabdf"
	fmt.Print(findString(a, b))
}

func findString(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	target := make(map[rune]int, len(s1))
	for _, c := range s1 {
		target[c]++
	}

	left, right := 0, len(s1)
	count := 0
	for right <= len(s2) {
		win := make(map[rune]int, len(s1))
		for i := left; i < right; i++ {
			win[rune(s2[i])]++
		}
		for key, value := range win {
			if target[key] == value {
				count++
			}
		}
		if count == len(s1) {
			return true
		} else {
			left++
			right++
		}
	}
	return false
}

// s1 = "ab" s2 = "rabv" true
