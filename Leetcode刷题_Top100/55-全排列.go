package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}                // 结果
	track := []int{}                // 路径
	used := make([]bool, len(nums)) // 决策否

	backtrack(nums, track, used, &res)
	return res
}

func backtrack(nums []int, track []int, used []bool, res *[][]int) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := range nums {
		// 排除不合法的选择
		if used[i] {
			continue
		}
		// 做选择
		track = append(track, nums[i])
		used[i] = true
		// 进入下一层决策树
		backtrack(nums, track, used, res)
		// 取消选择
		track = track[:len(track)-1]
		used[i] = false
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Print(permute(nums))
}
