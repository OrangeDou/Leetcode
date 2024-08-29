package main

import (
	"fmt"
	"sort"
)

// 1,1,2,2,3
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	//nums = sigle(nums)

	res := make([][]int, 0)
	track := make([]int, 0)

	var backTrace func(int)
	mapu := make(map[[3]int]bool, 0)
	backTrace = func(n int) {
		if len(track) == 3 && isZero(track) {
			res = append(res, append([]int(nil), track...))
			return
		}
		for i := n; i < len(nums); i++ {
			track = append(track, nums[i])
			backTrace(i + 1)
			track = track[:len(track)-1]
		}
	}
	backTrace(0)
	fmt.Print(res)
}

func isZero(nums []int) bool {
	if nums[0]+nums[1]+nums[2] == 0 {
		return true
	}
	return false
}

func sigle(nums []int) []int {
	mp := make(map[int]bool)
	new := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = true
	}
	for i, _ := range mp {
		new = append(new, i)
	}
	return new
}
