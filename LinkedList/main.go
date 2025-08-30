package main

func main() {
	list := DLL{}
	list.Insert(15)
	list.Insert(20)
	list.Insert(10)
	list.Insert(5)
	list.Display()
}

// func main() {
// 	list := LinkedList{}
// 	list.Insert(10)
// 	list.Insert(12)
// 	list.Insert(69)
// 	list.Insert(8)
// 	list.InsertLast(99)
// 	list.InsertAtIndex(100, 3)
// 	list.Display()

// 	// delete first
// 	fmt.Println("Delete first node...")
// 	list.DeleteFirst()
// 	fmt.Println("after delete the first node")

// 	list.Display()

// 	// delete last
// 	fmt.Println("Delete the last node")
// 	list.DeleteLast()
// 	fmt.Println("after delete the last node")

// 	list.Display()

// 	// delete at index
// 	fmt.Println("Delete the node at 2nd index which is 100")
// 	list.DeleteAt(2)
// 	fmt.Println("after deleting the node of index 2nd")

// 	list.Display()
// }
