package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	n          int = 3
	start, end time.Time
	a, b, c    string
)

func main() {
	//计时
	start := time.Now()
	defer func() {
		fmt.Print("\n")
		end := time.Now()
		fmt.Println("程序运行总共花费：", end.Second()-start.Second(), "秒")
	}()

	//交谈次数
	file, _ := os.ReadFile("test.txt")
	scanner := bufio.NewScanner(file)
	person0 := dialog[0]
	person1 := dialog[1]
	person2 := dialog[2]
	result := fellHapppy(person0, person1, person2)

	fmt.Printf("%d %d %d", result[0], result[1], result[2])

}
func fellHapppy(person0, person1, person2 string) []int {
	happy := make([]int, 3)
	for i := 0; i < n; i++ {
		if person0[i] != person1[i] && person0[i] != person2[i] && person1[i] != person2[i] {
			continue
		}
		if person0[i] == person1[i] && person0[i] == person2[i] && person1[i] == person2[i] {
			happy[0]++
			happy[1]++
			happy[2]++
		} else {
			if person0[i] == person1[i] {
				happy[0]++
				happy[1]++
			}
			if person0[i] == person2[i] {
				happy[0]++
				happy[2]++
			}
			if person1[i] == person2[i] {
				happy[1]++
				happy[2]++
			}
		}
	}
	return happy
}
