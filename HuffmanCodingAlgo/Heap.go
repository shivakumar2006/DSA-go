// heap

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
	if h.list[index] > h.list[p] {
		h.swap(index, p)
		h.upHeap(p)
	}
}

func (h *Heap[T]) Remove() (T, error) {
	var zero T
	if len(h.list) == 0 {
		return zero, errors.New("list is empty!")
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

	sorted, _ := h.HeapSort()
	fmt.Println("Sort : ", sorted)
}
