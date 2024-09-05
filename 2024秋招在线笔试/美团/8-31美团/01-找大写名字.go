package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	names := make([]string, 0)
	scanner := bufio.NewReader(os.Stdin)
	name, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(name)
	names = strings.Fields(name)
	// names = append(names, name...)
	fmt.Print(names)
}
