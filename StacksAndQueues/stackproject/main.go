package main

import (
	"fmt"
)

func main() {
	ds := NewDynamicStack(2)

	ds.Push(10)
	ds.Push(20)
	ds.Push(30)
	ds.Push(40) // automatically expands

	fmt.Println(ds.data)

	value, _ := ds.Pop()
	fmt.Println("Popped", value)
}
