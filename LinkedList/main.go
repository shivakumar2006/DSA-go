package main

func main() {
	list := LinkedList{}
	list.Insert(10)
	list.Insert(12)
	list.Insert(69)
	list.Insert(8)
	list.InsertLast(99)
	list.InsertAtIndex(100, 3)
	list.Display()
	list.DeleteFirst()
	list.Display()
	list.DeleteLast()
	list.Display()
}
