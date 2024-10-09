package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	leftmost := search(nums, target)
	if leftmost > 0 {
		if leftmost == len(nums) || nums[leftmost] != target {
			return []int{-1, -1}
		} else if nums[leftmost-1] == target {
			leftmost = search(nums[:leftmost-1], target)
		}
	} else if nums[leftmost] > target {
		leftmost = -1
	}
	rightmost := search(nums, target+1)
	if nums[rightmost-1] == target {
		rightmost = search(nums[:rightmost-1], target+1)
	}
	return []int{leftmost, rightmost}
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right + 1
}

func main() {
	a := []int{0, 0, 1, 2, 3, 3, 4}
	fmt.Print(searchRange(a, 2))
}
