package main

import "fmt"

type Stack struct {
	s     []interface{}
	limit int
}

func (s *Stack) Pop() interface{} {
	if len(s.s)-1 >= 0 {
		last := s.s[len(s.s)-1]
		s.s = s.s[:len(s.s)-1]
		return last
	} else {
		return nil
	}
}

func (s *Stack) Shift() interface{} {
	if len(s.s)-1 >= 0 {
		old := s.s[0]
		s.s = s.s[1:len(s.s)]
		return old
	} else {
		return nil
	}
}

func (s *Stack) Push(i interface{}) {
	s.s = append(s.s, i)
	if len(s.s) > s.limit {
		s.Shift()
	}
}

func (s *Stack) UpdateLimit(i int) {
	s.limit = i
}

func main() {
	s := &Stack{limit: 2}
	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")
	fmt.Println(s.Pop()) // -> "dataC"
	fmt.Println(s.Pop()) // -> "dataB"
	fmt.Println(s.Pop()) // -> nil
	s.UpdateLimit(3)
	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")
	fmt.Println(s.Pop()) // -> "dataC"
	fmt.Println(s.Pop()) // -> "dataB"
	fmt.Println(s.Pop()) // -> "dataA"
}
