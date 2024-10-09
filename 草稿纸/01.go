package main

import "fmt"

// func main() {

// 	a := 1
// 	b := 2
// 	swap2(a, b)
// 	swap1(&a, &b)
// 	fmt.Print(a, b)
// }

func swap1(a, b *int) {
	*a, *b = *b, *a
}
func swap2(a, b int) {
	a, b = b, a
}

func incr(a int) (b int) {
	defer func() {
		a++
		b++
	}()
	a++
	return a
}

func main() {
	var a, b int
	b = incr(a)
	fmt.Print(a, b)
}
