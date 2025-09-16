// // Segment Trees

// package main

// import "fmt"

// type SegmentTree struct {
// 	n    int
// 	tree []int
// }

// func NewSegmentTree(arr []int) *SegmentTree {
// 	n := len(arr)
// 	st := &SegmentTree{
// 		n:    n,
// 		tree: make([]int, 4*n),
// 	}
// 	st.build(arr, 0, 0, n-1)
// 	return st
// }

// func (st *SegmentTree) build(arr []int, node, start, end int) {
// 	if start == end {
// 		st.tree[node] = arr[start]
// 	} else {
// 		mid := (start + end) / 2
// 		st.build(arr, 2*node+1, start, mid)
// 		st.build(arr, 2*node+2, mid+1, end)
// 		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
// 	}
// }

// func (st *SegmentTree) Query(L, R int) int {
// 	return st.query(0, 0, st.n-1, L, R)
// }

// func (st *SegmentTree) query(node, start, end, L, R int) int {
// 	if R < start || end < L {
// 		return 0 // out of bound
// 	}
// 	if L <= start && end <= R {
// 		return st.tree[node]
// 	}
// 	mid := (start + end) / 2
// 	left := st.query(2*node+1, start, mid, L, R)
// 	right := st.query(2*node+2, mid+1, end, L, R)
// 	return left + right
// }

// func (st *SegmentTree) Update(index, value int) {
// 	st.update(0, 0, st.n-1, index, value)
// }

// func (st *SegmentTree) update(node, start, end, index, value int) {
// 	if start == end {
// 		st.tree[node] = value
// 	} else {
// 		mid := (start + end) / 2
// 		if index <= mid {
// 			st.update(2*node+1, start, mid, index, value)
// 		} else {
// 			st.update(2*node+2, mid+1, end, index, value)
// 		}
// 		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
// 	}
// }

func main() {
	arr := []int{1, 3, 5, 7, 11}
	st := NewSegmentTree(arr)

	fmt.Println("Original tree : ", arr)
	fmt.Println("Query(1, 3): ", st.Query(1, 3)) // 3+5+7 = 15

	st.Update(1, 10)
	fmt.Println("Update value arr[1] = 10 : ")
	fmt.Println("After update : ", st.Query(1, 3)) // 10+5+7 = 22
}

// package main

// import "fmt"

// type SegmentTree struct {
// 	n    int
// 	tree []int
// }

// func NewSegmentTree(arr []int) *SegmentTree {
// 	n := len(arr)
// 	st := &SegmentTree{
// 		n:    n,
// 		tree: make([]int, 4*n),
// 	}
// 	st.build(arr, 0, 0, n-1)
// 	return st
// }

// func (st *SegmentTree) build(arr []int, node, start, end int) {
// 	if start == end {
// 		st.tree[node] = arr[start]
// 	} else {
// 		mid := (start + end) / 2
// 		st.build(arr, 2*node+1, start, mid)
// 		st.build(arr, 2*node+2, mid+1, end)
// 		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
// 	}
// }

// func (st *SegmentTree) Query(L, R int) int {
// 	return st.query(0, 0, st.n-1, L, R)
// }

// func (st *SegmentTree) query(node, start, end, L, R int) int {
// 	if R < start || end < L {
// 		return 0 // out of bound
// 	}
// 	if L <= start && end <= R {
// 		return st.tree[node]
// 	}
// 	mid := (start + end) / 2
// 	left := st.query(2*node+1, start, mid, L, R)
// 	right := st.query(2*node+2, mid+1, end, L, R)

// 	return left + right
// }

// func (st *SegmentTree) Update(index, value int) {
// 	st.update(0, 0, st.n-1, index, value)
// }

// func (st *SegmentTree) update(node, start, end, index, value int) {
// 	if start == end {
// 		st.tree[node] = value
// 	} else {
// 		mid := (start + end) / 2
// 		if index <= mid {
// 			st.update(2*node+1, start, mid, index, value)
// 		} else {
// 			st.update(2*node+2, mid+1, end, index, value)
// 		}
// 		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
// 	}
// }

// func main() {
// 	arr := []int{1, 3, 5, 7, 11}
// 	st := NewSegmentTree(arr)

// 	fmt.Println("Original tree : ", arr)
// 	fmt.Println("Query(1, 3): ", st.Query(1, 3))

// 	st.Update(1, 10)
// 	fmt.Println("After update new value in arr[1] = 10: ")
// 	fmt.Println("10 at index 1 : ", st.Query(1, 3))
// }

// package main

// import "fmt"

// type SegmentTree struct {
// 	n    int
// 	tree []int
// }

// // constructor
// func NewSegmentTree(arr []int) *SegmentTree {
// 	n := len(arr)
// 	st := &SegmentTree{
// 		n:    n,
// 		tree: make([]int, 4*n),
// 	}
// 	st.build(arr, 0, 0, n-1)
// 	return st
// }

// func (st *SegmentTree) build(arr []int, node, start, end int) {
// 	if start == end {
// 		st.tree[node] = arr[start]
// 	} else {
// 		mid := (start + end) / 2
// 		st.build(arr, 2*node+1, start, mid)
// 		st.build(arr, 2*node+2, mid+1, end)
// 		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2] // parent sum the values of their child and stores it...
// 	}
// }

// // range query left and right
// func (st *SegmentTree) Query(L, R int) int {
// 	return st.query(0, 0, st.n-1, L, R)
// }

// func (st *SegmentTree) query(node, start, end, L, R int) int {
// 	if R < start || end < L {
// 		return 0 // out of range
// 	}
// 	if L <= start && end <= R {
// 		return st.tree[node]
// 	}
// 	mid := (start + end) / 2
// 	left := st.query(2*node+1, start, mid, L, R)
// 	right := st.query(2*node+2, mid+1, end, L, R)
// 	return left + right
// }

// // point update
// func (st *SegmentTree) Update(index, value int) {
// 	st.update(0, 0, st.n-1, index, value)
// }

// func (st *SegmentTree) update(node, start, end, index, value int) {
// 	if start == end {
// 		st.tree[node] = value
// 	} else {
// 		mid := (start + end) / 2
// 		if index <= mid {
// 			st.update(2*node+1, start, mid, index, value)
// 		} else {
// 			st.update(2*node+2, mid+1, end, index, value)
// 		}
// 		st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
// 	}
// }

// func main() {
// 	arr := []int{1, 3, 5, 7, 9, 11}
// 	st := NewSegmentTree(arr)

// 	fmt.Println("Original Array", arr)
// 	fmt.Println("Query(1, 3): ", st.Query(1, 3)) // sum of arr [1..3] = 3+5+7 = 15

// 	st.Update(1, 10) // arr[1] = 10
// 	fmt.Println("After update arr[1]=10")
// 	fmt.Println("Query(1,3):", st.Query(1, 3)) // sum of arr[1..3] = 10+5+7 = 22
// }
