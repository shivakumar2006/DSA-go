// // Huffman coder

package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	char  rune
	freq  int
	left  *Node
	right *Node
}

type priorityQueue []*Node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type HuffManCoder struct {
	encoder map[rune]string
	decoder map[string]rune
}

func NewHuffmanCoder(input string) *HuffManCoder {
	freqMap := make(map[rune]int)
	for _, ch := range input {
		freqMap[ch]++
	}

	// create min-heap
	pq := &priorityQueue{}
	heap.Init(pq)
	for ch, f := range freqMap {
		node := &Node{char: ch, freq: f}
		heap.Push(pq, node)
	}

	// build tree
	for pq.Len() > 1 {
		first := heap.Pop(pq).(*Node)
		second := heap.Pop(pq).(*Node)
		newNode := &Node{char: 0, freq: first.freq + second.freq, left: first, right: second}
		heap.Push(pq, newNode)
	}

	root := heap.Pop(pq).(*Node)

	hc := &HuffManCoder{
		encoder: make(map[rune]string),
		decoder: make(map[string]rune),
	}

	hc.initEncoderDecoder(root, "")
	return hc
}

func (hc *HuffManCoder) initEncoderDecoder(node *Node, path string) {
	if node == nil {
		return
	}

	if node.left == nil && node.right == nil {
		hc.encoder[node.char] = path
		hc.decoder[path] = node.char
	}

	hc.initEncoderDecoder(node.left, path+"0")
	hc.initEncoderDecoder(node.right, path+"1")
}

func (hc *HuffManCoder) Encode(source string) string {
	result := ""
	for _, ch := range source {
		result += hc.encoder[ch]
	}
	return result
}

func (hc *HuffManCoder) Decode(coded string) string {
	result := ""
	key := ""
	for _, bit := range coded {
		key += string(bit)
		if ch, ok := hc.decoder[key]; ok {
			result += string(ch)
			key = ""
		}
	}
	return result
}

func main() {
	input := "abracadabra"
	hc := NewHuffmanCoder(input)

	encoded := hc.Encode(input)
	decoded := hc.Decode(encoded)

	fmt.Println("Original:", input)
	fmt.Println("Encoded :", encoded)
	fmt.Println("Decoded :", decoded)
}

// package main

// import (
// 	"container/heap"
// 	"fmt"
// )

// type Node struct {
// 	char  rune
// 	freq  int
// 	left  *Node
// 	right *Node
// }

// // priority queue implements heap.interface for *Node
// type PriorityQueue []*Node

// func (pq PriorityQueue) Len() int {
// 	return len(pq)
// }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	return pq[i].freq < pq[j].freq
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }

// func (pq *PriorityQueue) Push(x interface{}) {
// 	*pq = append(*pq, x.(*Node))
// }

// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	*pq = old[0 : n-1]
// 	return item
// }

// // huffman code struct
// type HuffmanCoder struct {
// 	encoder map[rune]string
// 	decoder map[string]rune
// }

// // newHuffcoder builds encoder/decoder maps from the stirng
// func NewHuffmanCoder(input string) *HuffmanCoder {
// 	// frequency map
// 	freqMap := make(map[rune]int)
// 	for _, ch := range input {
// 		freqMap[ch]++
// 	}

// 	// creates min-heap
// 	pq := &PriorityQueue{}
// 	heap.Init(pq)
// 	for ch, f := range freqMap {
// 		node := &Node{char: ch, freq: f}
// 		heap.Push(pq, node)
// 	}

// 	// build tree
// 	for pq.Len() > 1 {
// 		first := heap.Pop(pq).(*Node)
// 		second := heap.Pop(pq).(*Node)
// 		newNode := &Node{char: 0, freq: first.freq + second.freq, left: first, right: second}
// 		heap.Push(pq, newNode)
// 	}

// 	// root of huffman tree
// 	root := heap.Pop(pq).(*Node)

// 	hc := &HuffmanCoder{
// 		encoder: make(map[rune]string),
// 		decoder: make(map[string]rune),
// 	}

// 	hc.initEncoderDecoder(root, "")

// 	return hc
// }

// func (hc *HuffmanCoder) initEncoderDecoder(node *Node, path string) {
// 	if node == nil {
// 		return
// 	}

// 	if node.left == nil && node.right == nil {
// 		hc.encoder[node.char] = path
// 		hc.decoder[path] = node.char
// 		return
// 	}

// 	hc.initEncoderDecoder(node.left, path+"0")
// 	hc.initEncoderDecoder(node.right, path+"1")
// }

// // encode returns encoding strings
// func (hc *HuffmanCoder) Encode(source string) string {
// 	result := ""
// 	for _, ch := range source {
// 		result += hc.encoder[ch]
// 	}
// 	return result
// }

// func (hc *HuffmanCoder) Decode(coded string) string {
// 	result := ""
// 	key := ""
// 	for _, bit := range coded {
// 		key += string(bit)
// 		if ch, ok := hc.decoder[key]; ok {
// 			result += string(ch)
// 			key = ""
// 		}
// 	}
// 	return result
// }

// func main() {
// 	input := "ABRACADABRA"
// 	hc := NewHuffmanCoder(input)

// 	encoded := hc.Encode(input)
// 	decoded := hc.Decode(encoded)

// 	fmt.Println("Original:", input)
// 	fmt.Println("Encoded :", encoded)
// 	fmt.Println("Decoded :", decoded)
// }
