/*
对于一个仅由左括号'('和右括号')'组成的字符串，小红想知道它的最长合法前缀的长度是多少。
对于某一个前缀，我们定义它是合法的，当且仅当该前缀满足以下条件：
存在一种拆分方案，可以将该前缀拆分为若干对匹配的括号'()'。例如'()','()()','(())'都是合法的，而')()(','))'是非法的。
特殊的，空串我们认为也是合法的。
*/
/*"(())))"*/

package main

import (
	"fmt"
)

type stack struct {
	Val []string
}

func main() {
	var n int //（和)的个数
	fmt.Scan(&n)
	var char string
	var prefix int //前缀长度
	fmt.Scan(&char)

	qh := make([]string, n) //括号数组 ['( ',')' ,')' ,')']
	for index, val := range char {
		qh[index] = string(val)
	}

	s := stack{}
	// 计算
	for i := 0; i < n; i++ {
		if qh[i] == "(" {
			s.push(qh[i])
		} else if qh[i] == ")" {
			if s.isEmpty() {
				continue
			}
			s.pop()
			prefix += 2
		}

	}
	fmt.Print(prefix)

}

func (s *stack) pop() (string, error) {
	if len(s.Val) == 0 {
		return "", fmt.Errorf("error")
	}
	index := len(s.Val) - 1
	element := s.Val[index]
	s.Val = s.Val[:index]
	return element, nil
}
func (s *stack) push(target string) {
	s.Val = append(s.Val, target)
}

func (s *stack) isEmpty() bool {
	if len(s.Val) == 0 {
		return true
	}
	return false
}
