package main

import (
	"fmt"
	"sort"
)

// func merge(intervals [][]int) [][]int {
// 	length := len(intervals)
// 	result := [][]int{}
// 	for i := 0; i < length-1; i++ {
// 		end := intervals[i][1]
// 		first := intervals[i+1][0]
// 		j := i

// 		if intervals[i][1] < intervals[i+1][1] {
// 			result = append(result, intervals[i])
// 		} else {
// 			tempArray := make([]int, 2)
// 			for first <= end && j <= length {
// 				tempArray = append(tempArray, intervals[i][0])
// 				j++
// 				end = intervals[j][1]
// 				first = intervals[j+1][0]

// 			}
// 			tempArray = append(tempArray, end)
// 			result = append(result, tempArray)
// 			i += j
// 		}

// 	}
// 	return result
// }

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {7, 10}}
	fmt.Print(merge(intervals))
}

// 官方题解：按照区间的左端点排序
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start, end := intervals[0][0], intervals[0][1]
	ans := [][]int{}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			ans = append(ans, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		} else {
			end = max(end, intervals[i][1])
		}
	}
	ans = append(ans, []int{start, end}) //追加末尾元素
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
