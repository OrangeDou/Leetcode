/*
多多携带一件价值为x的礼物参加了一个节日派对。
除多多外在场的有n个人，第i个人的当前持有的礼物价值为 ai
多多可以和任意当前持有礼物价值比多多低的人交换礼物。
请问最少经过多少次交换，可以使得n个人持有的礼物价值形成单调不减的数列
对于60%的数据，1<= n <= 10**3
对于100%的数据， 1<= n <= 2*10**6, x >= 1, ai >= 10**9
*/

package main

import "fmt"

// 我的思路
func main() {
	var n, x int64
	fmt.Scan(&n, &x)
	valueSet := make([]int64, 0)
	for i := int64(0); i < n; i++ {
		var ai int64
		fmt.Scan(&ai)
		valueSet = append(valueSet, ai)
	}
	result := helpFind(valueSet, x, n)
	fmt.Print(result)

}

func helpFind(valSet []int64, target int64, length int64) int64 {
	result := int64(0)
	if isUpper(valSet) {
		return 0
	} else {
		for i := length - 1; i >= 0; i-- {
			if valSet[i] < target {
				target, valSet[i] = valSet[i], target
				result++
			}
		}
		if isUpper(valSet) {
			return result
		} else {
			return -1
		}
	}

}
func isUpper(valSet []int64) bool {
	for i := 0; i < len(valSet)-1; i++ {
		if valSet[i+1] < valSet[i] {
			return false
		}
	}
	return true
}
