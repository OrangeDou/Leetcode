/*
a和b在玩一个游戏。有一个长度为n的数组m,他们玩q轮游戏，每轮游戏都是独立的。游戏规则如下，双方都会执行最优策略：
（1）第一步：游戏给出一个区间【l,r】
（2）第二部：a在【l,r】中选择一个数
（3）第三步：b将区间扩展成[L,R],其中[L,R]必须包含【l,r】，然后在【L，R】中选择一个数，但是不能跟a选的一样。
（4）第四步：a和b选择数字大的一方获胜，若相同则平局。
输入描述：第一行两个整数n,q（2<=n, q <= 2*10^5）表示数组长度和游戏次数，第二行输入n个整数t（1<= t <= 10^9 ），表示数组中的元素。接下来q行，每行两个整数，表示游戏次数。
输出描述：每次游戏先输出一行，若a获胜则输出“win”，平局输出“pin”，失败输出“lose”，再输出一行，表示达到最终状态所需的区间长度[L,R]的最小值。golang如何解决
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	m := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
	}
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		l--
		r--
		result, minPlan := play(m, l, r)
		switch result {
		case 1:
			fmt.Println("win")
		case 0:
			fmt.Println("pin")
		case -1:
			fmt.Println("lose")
		}
		fmt.Println(minPlan)
	}
}

func play(m []int, l, r int) (int, int) {
	maxA := -1
	indexA := -1
	for i := l; i <= r; i++ {
		if m[i] > maxA {
			maxA = m[i]
			indexA = i
		}
	}

	maxB := -1
	minLen := math.MaxInt32
	for L := 0; L <= l; L++ {
		for R := r; R < len(m); R++ {
			if L <= indexA && indexA <= R {
				for i := L; i <= R; i++ {
					if i != indexA && m[i] > maxB {
						maxB = m[i]
					}
				}
				if maxB > -1 {
					currLen := R - L + 1
					if currLen < minLen {
						minLen = currLen
					}
				}
			}
		}
	}
	if maxA > maxB {
		return 1, minLen
	} else if maxA < maxB {
		return -1, minLen
	} else {
		return 0, minLen
	}
}
