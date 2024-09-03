package main

import (
	"fmt"
)

func main() {
	var n, q int
	// n = 6
	// q = 2
	// executes := make([][]int, 0)
	// executes = append(executes, []int{1, 3})
	// executes = append(executes, []int{2, 4})
	fmt.Scan(&n, &q)
	executes := make([][]int, 0)
	for i := 0; i < q; i++ {
		var l, r int
		exectue := make([]int, 2)
		fmt.Scan(&l, &r)
		exectue[0] = l
		exectue[1] = r
		executes = append(executes, exectue)
	}
	// 构造数组
	nums := make([]int, 0)
	// 记录结果
	show := make(map[int][]int)
	for j := 1; j < n+1; j++ {
		nums = append(nums, j)
		show[j] = append(show[j], j)
	}
	// q次操作executes[k]
	for k := 0; k < q; k++ {
		left, right := executes[k][0], executes[k][1]
		nextArray := moveArray(nums, left, right)
		nums = nextArray
		for index := 0; index < len(nextArray); index++ {
			show[nextArray[index]] = append(show[nextArray[index]], index+1)
		}
	}
	for p := 1; p < n+1; p++ {
		times := couneTimes(show[p])
		fmt.Print(times)
		fmt.Print(" ")
	}

}

func moveArray(num []int, start, end int) []int {
	if start < 0 || end >= len(num) || start > end {
		return num
	}
	forword := num[start-1 : end]
	newNum := append([]int{}, num[:start-1]...)
	newNum = append(newNum, num[end:]...)
	newNum = append(newNum, forword...)
	return newNum
}

func couneTimes(array []int) int {

	countM := make(map[int]int)
	for _, v := range array {
		countM[v]++
	}
	return len(countM)
}
