package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var str string
	fmt.Scan(&str)
	var count int
	left, right := 0, 1
	for right < n && k > 0 {
		curStr := str[left:right]
		if curStr != "M" && curStr != "T" {
			count++
			k--
			left++
			right++
		} else {
			count++
			left++
			right++
		}
	}
	suspendStr := str[left:]
	for _, v := range suspendStr {
		if v == 'M' || v == 'T' {
			count++
		}
	}

	fmt.Print(count)
}
