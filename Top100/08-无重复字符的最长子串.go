package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	maxLength := 0

	for i := 0; i < len(s); i++ {
		currentLength := 0
		hashMap := make(map[string]bool, 0)
		for j := i; j < len(s); j++ {
			if hashMap[s[j:j+1]] {
				break
			}
			hashMap[s[j:j+1]] = true
			currentLength += 1
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}

	}
	return maxLength

}

func main() {
	s := "pwwkew"
	fmt.Print(lengthOfLongestSubstring2(s))
}

// 空间复杂度优化
func lengthOfLongestSubstring2(s string) int {
	tmpMap := map[byte]int{}
	start := 0
	result := 0
	for end := 0; end < len(s); end++ {
		if index, ok := tmpMap[s[end]]; ok && index+1 > start {
			start = index + 1
		}
		result = max(result, end-start+1)
		tmpMap[s[end]] = end
	}
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
