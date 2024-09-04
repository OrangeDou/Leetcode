package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	Value int
	body  []int
}

func main() {
	var n int
	fmt.Scan(&n)
	st := &stack{}
	output := make([]string, 0)
	for i := 0; i < n; i++ {
		var command string
		fmt.Scan(&command)
		switch command {
		case "push":
			var value int
			fmt.Scan(&value)
			st.Push(value)
		case "pop":
			res := st.Pop()
			re := strconv.Itoa(res)
			if res == -1 {
				output = append(output, "error")
			} else {
				output = append(output, re)
			}
		case "top":
			res := st.top()
			re := strconv.Itoa(res)
			if res == -1 {
				output = append(output, "error")
			} else {
				output = append(output, re)
			}
		}
	}
	for _, v := range output {
		fmt.Println(v)
	}
}

func (s *stack) Push(value int) {
	s.body = append(s.body, value)
}
func (s *stack) Pop() int {
	length := len(s.body)
	if length > 0 {
		num := s.body[length-1]
		s.body = s.body[:length-1]
		return num
	} else {
		return -1
	}
}

func (s *stack) top() int {
	if len(s.body) > 0 {
		return s.body[len(s.body)-1]
	} else {
		return -1
	}
}
