// stack exception

package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
	size int
}

// Push adds an element to the stack
func (s *Stack) Push(value int) error {
	if s.size >= 5 {
		return errors.New("Stack overflow")
	}
	s.data = append(s.data, value)
	s.size++
	return nil
}

// pop remove the top element
func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return 0, errors.New("Stack underflow")
	}
	top := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return top, nil
}

// peek return the top element without removing
func (s *Stack) Peek() (int, error) {
	if s.size == 0 {
		return 0, errors.New("Stack is empty.")
	}
	return s.data[s.size-1], nil
}

func main() {
	stack := &Stack{}

	// Push elements
	for i := 1; i <= 6; i++ {
		err := stack.Push(i * 10)
		if err != nil {
			fmt.Println("Error : ", err)
		} else {
			fmt.Println("Pushed : ", i*10)
		}
	}

	// pop elements
	for i := 1; i <= 6; i++ {
		value, err := stack.Pop()
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("Popped", value)
		}
	}
}
