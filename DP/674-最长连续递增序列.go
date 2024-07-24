package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 7}
	fmt.Print(findLengthOfLCIS2(nums))
}

// 暴力
func findLengthOfLCIS(nums []int) int {
	length := 1
	if len(nums) == 1 {
		return length
	}
	for i := 0; i < len(nums); i++ {
		curLength := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				curLength++
				i++
				continue
			} else {
				break
			}
		}
		if curLength > length {
			length = curLength
		}

	}
	return length

}

var maxLength int

// 优化 有问题
func findLengthOfLCIS2(nums []int) int {
	hash := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if hash[i] < maxLength {
			break
		}
		help(i, nums, hash)
	}
	return maxLength

}
func help(index int, nums []int, mp map[int]int) {
	length := 0

	for i := index; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			help(i+1, nums[i:], mp)
			length++
		}
		mp[index] = length
		if length > maxLength {
			maxLength = length
		}
	}
}
