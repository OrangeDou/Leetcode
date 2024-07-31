package main

func removeElement(nums []int, val int) int {
	f, s := 0, 0
	for f < len(nums) {
		if nums[f] != val {
			nums[s] = nums[f]
			s++
		}
		f++
	}
	return s
}
