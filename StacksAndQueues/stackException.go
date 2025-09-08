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
}
