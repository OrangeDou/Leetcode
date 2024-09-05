package main

import (
	"fmt"
	"strconv"
)

func findAllSubarrays(arr []int, k int) int {
	n := len(arr)
	res := make([][]int, 0)
	count := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			subarray := arr[i : j+1]
			res = append(res, subarray)
		}
	}
	for i := 0; i < len(res); i++ {
		if deleteAndCalculate(arr, res[i], k) {
			count++
		}
	}
	return count
}

func main() {
	arr := []int{2, 5, 3, 4, 20}

	result := findAllSubarrays(arr, 2)
	fmt.Print(result)
}
func deleteAndCalculate(arr []int, subarray []int, k int) bool {
	var curArray []int = make([]int, len(arr))
	copy(curArray, arr)
	zero := "0"
	for o := 1; o < k; {
		zero += "0"
		o++
	}
	// delete
	for i := 0; i < len(subarray); i++ {
		for j := 0; j < len(arr); j++ {
			if curArray[j] == subarray[i] {
				curArray[j] = 1
				break
			}
		}
	}
	// calculate
	var count int = 1
	for i := 0; i < len(curArray); i++ {
		count *= curArray[i]
	}
	numStr := strconv.Itoa(count)
	if len(numStr) < k {
		return false
	}
	if numStr[len(numStr)-k:] == zero {
		return true
	}
	return false
}
