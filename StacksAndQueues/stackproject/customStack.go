package main

import "errors"

type CustomStack struct {
	data []int
	size int
}

func NewCustomStack(capacity int) *CustomStack {
	return &CustomStack{data: []int{}, size: capacity}
}

func (s *CustomStack) Push(value int) error {
	if len(s.data) >= s.size {
		return errors.New("Stack is full!")
	}
	s.data = append(s.data, value)
	return nil
}

func (s *CustomStack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("Stack is empty")
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, nil
}
