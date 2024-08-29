package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var javaCode string
	emptyLineCount := 0 // 用于跟踪空行的数量

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			emptyLineCount++ // 当前行为空，增加空行计数
			if emptyLineCount == 2 {
				// 如果已经连续输入了两个空行，则结束循环
				break
			}
		} else {
			// 如果当前行不为空，则重置空行计数
			emptyLineCount = 0
		}
		javaCode += line + "\n"
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "读取输入时发生错误:", err)
		return
	}

	left, right := 0, 0
	right = left + 3
	count := 0
	for right < len(javaCode) {
		if javaCode[left:right] == "for" {
			count++

		}
		left++
		right = left + 3
	}
	fmt.Print(count)
}
