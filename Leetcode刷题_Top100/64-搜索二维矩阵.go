package main

func searchMatrix(matrix [][]int, target int) bool {
	rangeArray := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		rangeArray = append(rangeArray, matrix[i][0])
	}
	tarLine := binaryFind(rangeArray, target)
	if tarLine < 0 {
		return false
	}
	index := binaryFind(matrix[tarLine], target)
	if matrix[tarLine][index] == target {
		return true
	}
	return false
}

func binaryFind(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
