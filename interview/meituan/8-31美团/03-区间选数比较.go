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
