package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	var track []int

	var backtrack func(int)
	backtrack = func(start int) {
		res = append(res, append([]int(nil), track...))
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}

func main() {
	nums := [4]int{1, 2, 3, 4}
	fmt.Print(subsets(nums[:]))
}
