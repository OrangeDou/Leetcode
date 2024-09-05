package main

func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 1 {
		return nil
	}
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[right] + numbers[left]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}

}
