/*
输入：[1,5,2,4,3]
输出：3
解释：[1,2,4]
*/
package main

import "fmt"

func main() {
	nums := []int{1, 5, 2, 4, 3}
	fmt.Print(findSeq3(nums))
}

// 暴力
func findSeq(nums []int) int {
	length := 0
	for i := 0; i < len(nums); i++ {
		length = max(length, seqLength(nums, i))
	}
	return length

}

func seqLength(nums []int, i int) int {
	if i == len(nums)-1 {
		return 1
	}
	Length := 1
	for j := i + 1; j < len(nums); j++ {
		if nums[j] > nums[i] {
			Length = max(Length, seqLength(nums, j)+1)
		}
	}
	return Length
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// DP
func findSeq2(nums []int) int {
	length := 0
	for i := 0; i < len(nums); i++ {
		length = max(length, DP(nums, i))
	}
	return length

}
func DP(nums []int, i int) int {
	if i == len(nums)-1 {
		return 1
	}
	hashMap := make(map[int]int, 0)
	if hashMap[i] != 0 {
		return hashMap[i]
	}
	Length := 1
	for j := i + 1; j < len(nums); j++ {
		if nums[j] > nums[i] {
			Length = max(Length, seqLength(nums, j)+1)

		}
	}
	hashMap[i] = Length
	return Length
}

// 迭代

func findSeq3(nums []int) int {
	n := len(nums)
	result := make([]int, n)
	for i := n; i > 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				result[i] = max(result[i], result[j]+1)
			}
		}
	}
	return maxArray(result)
}

func maxArray(arr []int) int {
	max := 0
	for k := 0; k < len(arr); k++ {
		if arr[k] > max {
			max = arr[k]
		}
	}
	return max
}
