package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 函数用于奇妙约分
func wonderfulReduce(numerator, denominator string) (string, string) {
	// 从分子和分母中删除相同的数字
	for len(numerator) > 0 && len(denominator) > 0 {
		for i := 0; i < len(numerator); i++ {
			for j := 0; j < len(denominator); j++ {
				if numerator[i] == denominator[j] {
					// 找到相同的数字，删除它们
					numerator = numerator[:i] + numerator[i+1:] + numerator[i:][len(numerator)-1]
					denominator = denominator[:j] + denominator[j+1:] + denominator[j:][len(denominator)-1]
					goto Continue
				}
			}
		}
		// 如果没有找到相同的数字，退出循环
		break
	Continue:
	}

	// 处理前导0
	numerator = strings.TrimLeft(numerator, "0")
	denominator = strings.TrimLeft(denominator, "0")

	// 返回奇妙约分后的结果
	return numerator, denominator
}

func main() {
	scanner := bufio.Scanner{bufio.NewReader(os.Stdin)}

	// 读取行数
	fmt.Print("请输入行数 n: ")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 读取每一行的两个分数并进行处理
	for i := 0; i < n; i++ {
		fmt.Print("请输入两个分数（用空格分隔）: ")
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		firstFraction := parts[0]
		secondFraction := parts[1]

		// 分割分子和分母
		numerator1, denominator1 := "", ""
		numerator2, denominator2 := "", ""

		if strings.Contains(firstFraction, "/") {
			numerator1, denominator1 = strings.Split(firstFraction, "/")
		}
		if strings.Contains(secondFraction, "/") {
			numerator2, denominator2 = strings.Split(secondFraction, "/")
		}

		// 进行奇妙约分
		reducedNumerator1, reducedDenominator1 := wonderfulReduce(numerator1, denominator1)
		reducedNumerator2, reducedDenominator2 := wonderfulReduce(numerator2, denominator2)

		// 判断第二个分数是否等于奇妙约分后的第一个分数
		if reducedNumerator1 == reducedNumerator2 && reducedDenominator1 == reducedDenominator2 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
