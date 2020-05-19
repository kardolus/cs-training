package main

import (
	"fmt"
)

func main() {
	var (
		stack = NewStack(2)
		value int
		err   error
	)

	fmt.Println("Empty()", stack.Empty())
	fmt.Println("Size()", stack.Size())
	fmt.Println("Contains(3)", stack.Contains(3))

	fmt.Println("----> Push(3)")
	stack.Push(3)
	fmt.Println("Empty()", stack.Empty())
	fmt.Println("Size()", stack.Size())
	fmt.Println("Contains(3)", stack.Contains(3))
	fmt.Println("Contains(4)", stack.Contains(4))

	fmt.Println("----> Push(4)")
	stack.Push(4)
	fmt.Println("Empty()", stack.Empty())
	fmt.Println("Size()", stack.Size())
	fmt.Println("Contains(3)", stack.Contains(3))
	fmt.Println("Contains(4)", stack.Contains(4))

	value, err = stack.Pop()
	fmt.Println("Pop() value / err:", value, "/", err)
	fmt.Println("Size()", stack.Size())
	fmt.Println("Contains(4)", stack.Contains(4))

	value, err = stack.Pop()
	fmt.Println("Pop() value / err:", value, "/", err)
	fmt.Println("Empty()", stack.Empty())
	fmt.Println("Size()", stack.Size())
	fmt.Println("Contains(3)", stack.Contains(3))

	value, err = stack.Pop()
	fmt.Println("Pop() value / err:", value, "/", err)

	fmt.Println("----> Push(2)")
	stack.Push(2)
	fmt.Println("----> Push(6)")
	stack.Push(6)
	fmt.Println("----> Push(2)")
	stack.Push(2)
	fmt.Println("----> Push(9)")
	stack.Push(9)
	fmt.Println("Size()", stack.Size())

	value, err = stack.Peek()
	fmt.Println("Peek() value / err:", value, "/", err)
}

/*
  Implementation below
*/

type Stack struct {
	size int
	data []int
}

func NewStack(initialSize int) Stack {
	return Stack{
		size: 0,
		data: make([]int, initialSize),
	}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Contains(value int) bool {
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			return true
		}
	}

	return false
}

func (s *Stack) Push(value int) {
	s.copy()
	s.data[s.size] = value
	s.size++
}

func (s *Stack) Pop() (int, error) {
	result, err := s.Peek()
	if err != nil {
		return 0, err
	}
	s.size--

	return result, nil
}

func (s *Stack) Peek() (int, error) {
	if s.size == 0 {
		return 0, fmt.Errorf("empty stack exception\n")
	}

	return s.data[s.size-1], nil
}

func (s *Stack) copy() {
	if s.size == len(s.data) {
		newData := make([]int, s.size*2)

		copy(newData, s.data)
		s.data = newData
	}
}
