/*
小红和小紫正在玩一个游戏，每一关都有一个分数。
如果某人某一关分数比上一关高，但另一个人这一关分数比上一关低，那么他就可以嘲笑对方。
如果两个人这一关游戏的分数都比上一关多，则增量更多的可以嘲笑对方；
如果两个人这一关游戏的分数都比上一关少，则减量更少的可以嘲笑对方。只有当他们的增量相同或者减量相同时，才不会互相嘲笑。
例如，假设小红第一关的分数为5，第二关的关卡分数为10；小紫第一关的分数为2，第二关的分数为8，显然小红增加的比小紫多，那么小红就可以嘲笑小紫。
现在给定了小红和小紫每一关的分数，你可以选择一段连续的关卡，使得这一段关卡中两个人都不会互相嘲笑，问最多可以选择多少个关卡。特别的，一段连续关卡中的第一关两人不会互相嘲笑。
[]
[101011011010]
5
[1 2 3 4 5]
[-1 2 4 3 2]
*/
package main

import "fmt"

// 超时了
func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 0)
	b := make([]int, 0)
	for i := 0; i < n; i++ {
		var score int
		fmt.Scan(&score)
		a = append(a, score)
	}
	for i := 0; i < n; i++ {
		var score int
		fmt.Scan(&score)
		b = append(b, score)
	}
	result := make([]int, 0)
	result = append(result, 1)
	length := 0
	maxlength := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			length += 1
			continue
		}
		curA := a[i] - a[i-1]
		curB := b[i] - b[i-1]
		if a[i] > a[i-1] && b[i] > b[i-1] && curA == curB && length != 0 {
			length += 1
			if length > maxlength {
				maxlength = length
			}
			continue
		} else {
			length = 0
		}
		if curA < 0 && curB < 0 && curA == curB && length != 0 {
			length += 1
			if length > maxlength {
				maxlength = length
			}
			continue
		} else {
			length = 0
		}

	}
	fmt.Print(maxlength)
}
