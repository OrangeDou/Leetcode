package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	// 询问
	ask := make([][]int, q)
	for i := 0; i < q; i++ {
		ask[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			fmt.Scan(&ask[i][j])
		}
	}

	result := help(arr, ask)
	for _, v := range result {
		fmt.Println(v[0], v[1])
	}

}

func help(arr []int, ask [][]int) [][]int {
	var smallest, bigest int
	var result [][]int
	// 寻找未知元素的个数
	countZero, count := 0, 0
	for _, v := range arr {
		if v == 0 {
			countZero++
		}
		count += v
	}
	// 遍历询问，计算最小值和最大值
	for j := 0; j < len(ask); j++ {
		if ask[j][0] == ask[j][1] {
			smallest = count + countZero*ask[j][0]
			bigest = smallest
			result = append(result, []int{smallest, bigest})
		} else {
			smallest = count + countZero*ask[j][0]
			bigest = count + countZero*ask[j][1]
			result = append(result, []int{smallest, bigest})
		}
	}
	return result
}
