package main

import "fmt"

func findSubString(s string) []string {
	var result []string
	mpSubString := make(map[string]bool, 0)
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			if !mpSubString[s[i:j]] {
				mpSubString[s[i:j]] = true
				result = append(result, s[i:j])
			}
		}
	}
	return result
}

func main() {
	test := "hello"
	fmt.Print(findSubString(test))
}
