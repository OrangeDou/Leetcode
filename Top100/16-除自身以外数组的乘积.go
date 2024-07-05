package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Print(productExceptSelf2(nums))
}

// 超出时间范围
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result := 1
		arr := make([]int, len(nums))
		_ = copy(arr[:], nums[:])
		arr = append(arr[:i], arr[i+1:]...)
		for _, v := range arr {
			result = v * result
		}
		answer[i] = result
	}
	return answer
}

//O（n）左右乘积列表

func productExceptSelf2(nums []int) []int {
	pro := make([]int, len(nums))
	end := make([]int, len(nums))
	//result := make([]int, len(nums))
	pro[0] = 1
	end[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		pro[i] = pro[i-1] * nums[i-1]
	}
	for j := len(nums) - 2; j >= 0; j-- {
		end[j] = end[j+1] * nums[j+1]
	}

	for k := 0; k < len(nums); k++ {
		nums[k] = pro[k] * end[k]
	}
	return nums
}
