package main

import "fmt"

type Stack struct {
	data []any
	top  int
}

func (s *Stack) Push(val any) {
	s.data = append(s.data, val)
	s.top++
}

func (s *Stack) Pop() any {
	if len(s.data) <= 0 {
		return nil
	}

	s.top--
	val := s.data[s.top]
	s.data = s.data[:s.top]
	return val
}

func (s *Stack) Length() int {
	return len(s.data)
}

func maxDepth(s string) int {
	stack := Stack{}
	maxNested := 0

	for _, c := range s {
		if c != '(' && c != ')' {
			continue
		}

		if c == '(' {
			stack.Push(c)
			maxNested = max(maxNested, stack.Length())
		} else {
			stack.Pop()
		}
	}

	return maxNested
}

func maxDepthTest() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1")) // 3
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))   // 3
	fmt.Println(maxDepth("1+(2*3)/(2-1)"))       // 1
	fmt.Println(maxDepth("1"))                   // 0
}

// func main() {
// 	maxDepthTest()
// }
