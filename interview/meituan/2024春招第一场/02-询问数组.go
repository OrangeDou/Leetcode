package main

import (
	"bufio"
	"fmt"
	"os"
)

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

// 牛友思路：
const N int = 1e5 + 10

func main2() {
	in := bufio.NewReader(os.Stdin)
	n := 1
	q := 1
	fmt.Fscan(in, &n, &q)

	var s int64 = 0
	var zeroNum int64 = 0
	for i := 1; i <= n; i++ {
		var a int64
		fmt.Fscan(in, &a)
		if a == 0 {
			zeroNum++
		}
		s = s + a
	}
	for i := 0; i < q; i++ {
		var l, r int64
		fmt.Fscan(in, &l, &r)
		fmt.Printf("%d %d\n", s+zeroNum*l, s+zeroNum*r)
	}
}
