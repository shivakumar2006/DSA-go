package main

import "fmt"

func main() {
	dq := NewdynamicQueue(3)

	_ = dq.Push(10)
	_ = dq.Push(20)
	_ = dq.Push(30)
	fmt.Println("Queue : ", dq.data)

	// capacity will expand automatically here
	_ = dq.Push(40)
	_ = dq.Push(50)

	fmt.Println("Queue after dynamic push : ", dq.data)

}
