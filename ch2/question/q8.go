package main

import "fmt"

func main() {
	var s stack
	for i := 0; i < 12; i++ {
		s.push(i + 5)
	}
	s.pop()

	fmt.Printf("stack %v\n", s)
}

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(v int) {
	if s.i > 9 {
		return
	}
	s.data[s.i] = v
	s.i++
}

func (s *stack) pop() (v int) {
	if s.i < 0 {
		return
	}
	s.i--
	v = s.data[s.i]
	return
}
