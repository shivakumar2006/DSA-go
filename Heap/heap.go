// // Heap

package main

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	list []T
}

func NewHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{list: []T{}}
}

func (h *Heap[T]) swap(i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return (2 * index) + 1
}

func rightChild(index int) int {
	return (2 * index) + 2
}

func (h *Heap[T]) Insert(value T) {
	h.list = append(h.list, value)
	h.upHeap(len(h.list) - 1)
}

func (h *Heap[T]) upHeap(index int) {
	if index == 0 {
		return
	}

	p := parent(index)
	if h.list[index] < h.list[p] {
		h.swap(index, p)
		h.upHeap(p)
	}
}

func (h *Heap[T]) Remove() (T, error) {
	var zero T
	if len(h.list) == 0 {
		return zero, errors.New("removing the smallest element")
	}

	temp := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	if len(h.list) > 0 {
		h.list[0] = last
		h.downHeap(0)
	}
	return temp, nil
}

func (h *Heap[T]) downHeap(index int) {
	min := index
	left := leftChild(index)
	right := rightChild(index)

	if left < len(h.list) && h.list[min] > h.list[left] {
		min = left
	}
	if right < len(h.list) && h.list[min] > h.list[right] {
		min = right
	}
	if min != index {
		h.swap(min, index)
		h.downHeap(min)
	}
}

func (h *Heap[T]) HeapSort() ([]T, error) {
	var data []T
	for len(h.list) > 0 {
		item, err := h.Remove()
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func main() {
	h := NewHeap[int]()
	h.Insert(5)
	h.Insert(3)
	h.Insert(10)
	h.Insert(1)

	fmt.Println("Heap list : ", h.list)

	remove, _ := h.Remove()
	fmt.Println("Removed : ", remove)

	sorting, _ := h.HeapSort()
	fmt.Println("heap sort: ", sorting)
}

// package main

// import (
// 	"errors"
// 	"fmt"

// 	"golang.org/x/exp/constraints"
// )

// type Heap[T constraints.Ordered] struct {
// 	list []T
// }

// func NewHeap[T constraints.Ordered]() *Heap[T] {
// 	return &Heap[T]{list: []T{}}
// }

// func (h *Heap[T]) swap(i, j int) {
// 	h.list[i], h.list[j] = h.list[j], h.list[i]
// }

// func parent(index int) int {
// 	return (index - 1) / 2
// }

// func leftChild(index int) int {
// 	return (2 * index) + 1
// }

// func rightChild(index int) int {
// 	return (2 * index) + 2
// }

// func (h *Heap[T]) Insert(value T) {
// 	h.list = append(h.list, value)
// 	h.upHeap(len(h.list) - 1)
// }

// func (h *Heap[T]) upHeap(index int) {
// 	if index == 0 {
// 		return
// 	}

// 	p := parent(index)
// 	if h.list[index] < h.list[p] {
// 		h.swap(index, p)
// 		h.upHeap(p)
// 	}
// }

// // now, remove the smallest element
// func (h *Heap[T]) Remove() (T, error) {
// 	var zero T
// 	if len(h.list) == 0 {
// 		return zero, errors.New("Removing the smallest element!")
// 	}
// 	temp := h.list[0]
// 	last := h.list[len(h.list)-1]
// 	h.list = h.list[:len(h.list)-1]

// 	if len(h.list) > 0 {
// 		h.list[0] = last
// 		h.downHeap(0)
// 	}
// 	return temp, nil
// }

// func (h *Heap[T]) downHeap(index int) {
// 	min := index
// 	left := leftChild(index)
// 	right := rightChild(index)

// 	if left < len(h.list) && h.list[min] > h.list[left] {
// 		min = left
// 	}
// 	if right < len(h.list) && h.list[min] > h.list[right] {
// 		min = right
// 	}
// 	if min != index {
// 		h.swap(min, index)
// 		h.downHeap(min)
// 	}
// }

// func (h *Heap[T]) HeapSort() ([]T, error) {
// 	var data []T
// 	for len(h.list) > 0 {
// 		item, err := h.Remove()
// 		if err != nil {
// 			return nil, err
// 		}
// 		data = append(data, item)
// 	}
// 	return data, nil
// }

// func main() {
// 	h := NewHeap[int]()
// 	h.Insert(5)
// 	h.Insert(3)
// 	h.Insert(10)
// 	h.Insert(1)

// 	fmt.Println("Heap array : ", h.list)

// 	val, _ := h.Remove()
// 	fmt.Println("Removed: ", val)

// 	sort, _ := h.HeapSort()
// 	fmt.Println("Heap sort :", sort)
// }

// package main

// import (
// 	"errors"
// 	"fmt"

// 	"golang.org/x/exp/constraints"
// )

// // heap is generic min-heap
// type Heap[T constraints.Ordered] struct {
// 	list []T
// }

// // creates an empty heap
// func NewHeap[T constraints.Ordered]() *Heap[T] {
// 	return &Heap[T]{list: []T{}}
// }

// func (h *Heap[T]) swap(i, j int) {
// 	h.list[i], h.list[j] = h.list[j], h.list[i]
// }

// func parent(i int) int {
// 	return (i - 1) / 2
// }

// func leftChild(i int) int {
// 	return (2 * i) + 1
// }

// func rightChild(i int) int {
// 	return (2 * i) + 2
// }

// // insert pushes a value into a heap
// func (h *Heap[T]) Insert(value T) {
// 	h.list = append(h.list, value)
// 	h.upHeap(len(h.list) - 1)
// }

// func (h *Heap[T]) upHeap(i int) {
// 	if i == 0 {
// 		return
// 	}
// 	p := parent(i)
// 	if h.list[i] < h.list[p] {
// 		h.swap(i, p)
// 		h.upHeap(p)
// 	}
// }

// // remove pops the smallest element
// func (h *Heap[T]) Remove() (T, error) {
// 	var zero T
// 	if len(h.list) == 0 {
// 		return zero, errors.New("removing from an empty heap!")
// 	}

// 	temp := h.list[0]
// 	last := h.list[len(h.list)-1]
// 	h.list = h.list[:len(h.list)-1]

// 	if len(h.list) > 0 {
// 		h.list[0] = last
// 		h.downHeap(0)
// 	}
// 	return temp, nil
// }

// func (h *Heap[T]) downHeap(i int) {
// 	min := i
// 	l := leftChild(i)
// 	r := rightChild(i)

// 	if l < len(h.list) && h.list[min] > h.list[l] {
// 		min = l
// 	}
// 	if r < len(h.list) && h.list[min] > h.list[r] {
// 		min = r
// 	}
// 	if min != i {
// 		h.swap(min, i)
// 		h.downHeap(min)
// 	}
// }

// // heapsort remove all items and reutrns them in order.
// func (h *Heap[T]) HeapSort() ([]T, error) {
// 	var data []T
// 	for len(h.list) > 0 {
// 		item, err := h.Remove()
// 		if err != nil {
// 			return nil, err
// 		}
// 		data = append(data, item)
// 	}
// 	return data, nil
// }

// func main() {
// 	h := NewHeap[int]()
// 	h.Insert(5)
// 	h.Insert(3)
// 	h.Insert(10)
// 	h.Insert(1)

// 	fmt.Println("Heap array : ", h.list)

// 	val, _ := h.Remove()
// 	fmt.Println("Removed: ", val)

// 	sorted, _ := h.HeapSort()
// 	fmt.Println("Sorted : ", sorted)
// }
