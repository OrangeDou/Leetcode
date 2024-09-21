package main

import "fmt"

func main() {
	intex := [][]int{
		{12, 32, 9, 11, 34},
		{8, 54, 76, 23, 7},
		{27, 18, 25, 9, 43},
		{11, 23, 78, 63, 19},
		{9, 22, 56, 31, 5},
	}
	help(intex)

}
func help(arr [][]int) {
	if len(arr) == 0 {
		return
	}
	top, bottom := 0, len(arr)-1
	left, right := 0, len(arr)-1
	sum := 0

	for top <= bottom && left <= right {
		// 左到右
		for i := left; i <= right; i++ {
			if i == right {
				fmt.Printf("%d", arr[top][i])
				sum += arr[top][i]
			} else {
				fmt.Printf("%d+", arr[top][i])
				sum += arr[top][i]
			}

		}
		top++
		for i := top; i <= bottom; i++ {
			fmt.Printf("%d+", arr[i][right])
			sum += arr[i][right]
		}
		right--

		// 右到左
		if top <= bottom {
			for i := right; i >= left; i-- {
				fmt.Printf("%d+", arr[bottom][i])
				sum += arr[bottom][i]
			}
			bottom--
		}
		if left <= right {

			for i := bottom; i >= top; i-- {
				fmt.Printf("%d+", arr[i][left])
				sum += arr[i][left]
			}
			left++
		}
	}

	fmt.Printf("=%d", sum)
}
