package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	targetSet := make([]int, t)

	for i := 0; i < t; i++ {
		var count int
		mp := make(map[string]int)
		fmt.Scan(&count)
		for j := 0; j < count; j++ {
			var curStr string
			fmt.Scan(&curStr)
			mp[curStr]++
		}
		targetSet[i] = getKeyNum(mp)
	}
	for _, v := range targetSet {
		fmt.Println(v)
	}
}
func getKeyNum(mp map[string]int) (res int) {
	for k, v := range mp {
		if k[0:3] == "bnu" {
			res += v
		}
	}
	return
}
