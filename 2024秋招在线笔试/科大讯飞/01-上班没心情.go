package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	if t < 1 || t > 100 {
		return
	}
	work := make([][]int, 0)
	for i := 0; i < t; i++ {
		today := make([]int, 2)
		var A, B int
		fmt.Scan(&A, &B)
		today[0] = A
		today[1] = B
		work = append(work, today)
	}

	for j := 0; j < t; j++ {
		fmt.Println(workWithPress(work[j]))
	}

}

func workWithPress(today []int) int {
	fell := 0
	press := today[0]
	day := today[1]
	for i := 0; i < day; i++ {
		fell += press
	}
	return fell + day

}
