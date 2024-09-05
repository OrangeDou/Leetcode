package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	lineArray := strings.Split(line, " ")
	nStr := lineArray[0]
	kStr := lineArray[1]
	n, _ := strconv.Atoi(nStr)
	k, _ := strconv.Atoi(kStr)
	scanner.Scan()
	line2 := scanner.Text()
	lineArray2 := strings.Split(line2, " ")
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(lineArray2[i])
	}

}
