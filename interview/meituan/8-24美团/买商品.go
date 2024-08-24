/*
5 5
2 3 4 5 6
1 1 0 1 1
1 3 1 2
1 3 1 2
1 3 0 5
1 5 1 1
1 5 0 1
*/
package main

import "fmt"

type shangpin struct {
	types     int
	validDate int
	has       bool
}

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	date := make([]int, m)
	typeu := make([]int, m)
	people := make([][]int, 0)
	for i := 0; i < m; i++ {
		fmt.Scan(&date[i])
	}
	for j := 0; j < m; j++ {
		fmt.Scan(&typeu[j])
	}
	for k := 0; k < n; k++ {
		one := make([]int, 0)
		var l, r, target, num int
		fmt.Scan(&l, &r, &target, &num)
		one = append(one, l)
		one = append(one, r)
		one = append(one, target)
		one = append(one, num)
		people = append(people, one)
	}

	shangPinArr := make([]shangpin, 0)
	for o := 0; o < m; o++ {
		curShangPin := shangpin{
			types:     typeu[o],
			validDate: date[o],
			has:       true,
		}
		shangPinArr = append(shangPinArr, curShangPin)
	}

	for h := 0; h < n; h++ {
		res = shopping(people, shangPinArr)

	}
}

func shopping(peopleA []int, shangpinA []shangpin) []int {
	l, r := peopleA[0], peopleA[1]
	targetA := peopleA[2]
	wantNum := peopleA[3]
	curNum := 0
	win := shangpinA[l-1 : r]
	res := make([]int, 0)
	for i := 0; i < wantNum; i++ {
		for j := l - 1; j < len(win); j++ {
			if shangpinA[j].has && shangpinA[j].types == targetA {
				curNum++
				shangpinA[j].has = false
				res = append(res, curNum)
			} else {
				res = append(res, -1)
			}
		}
	}
	return res
}
