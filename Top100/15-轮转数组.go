package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Print(nums)
}

// 没有返回值 所有需要原地改动
func rotate(nums []int, k int) {
	n := len(nums)
	if k > n {
		k = k % n
	}
	result := make([]int, 0)
	for i := n - k; i < n; i++ {
		result = append(result, nums[i])
	}
	for j := 0; j < n-k; j++ {
		result = append(result, nums[j])
	}
	for k := 0; k < n; k++ {
		nums[k] = result[k]
	}

}

//官方题解
