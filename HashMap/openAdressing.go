package main

import "fmt"

type entry struct {
	key     string
	value   int
	used    bool // slot ever used
	deleted bool // tombstone marker
}

type HashMap struct {
	table      []*entry
	size       int
	loadFactor float64
}

func NewHashMap(capacity int) *HashMap {
	return &HashMap{
		table:      make([]*entry, capacity),
		loadFactor: 0.7, // resize when >70% full
	}
}

// simple hash function
func (h *HashMap) hash(key string) int {
	hash := 0
	for _, c := range key {
		hash += int(c)
	}
	return hash % len(h.table)
}

// resize & reinsert everything
func (h *HashMap) rehash() {
	old := h.table
	h.table = make([]*entry, len(old)*2)
	h.size = 0
	for _, e := range old {
		if e != nil && e.used && !e.deleted {
			h.Put(e.key, e.value)
		}
	}
}

// insert/update key
func (h *HashMap) Put(key string, value int) {
	if float64(h.size)/float64(len(h.table)) > h.loadFactor {
		h.rehash()
	}

	index := h.hash(key)
	start := index
	var firstTombstone int = -1

	for {
		e := h.table[index]

		if e == nil { // empty slot
			if firstTombstone != -1 {
				index = firstTombstone
			}
			h.table[index] = &entry{key: key, value: value, used: true}
			h.size++
			return
		}

		if e.used && !e.deleted && e.key == key { // update
			e.value = value
			return
		}

		if e.deleted && firstTombstone == -1 { // remember first tombstone
			firstTombstone = index
		}

		index = (index + 1) % len(h.table)
		if index == start { // full
			fmt.Println("HashMap full")
			return
		}
	}
}

// get value by key
func (h *HashMap) Get(key string) (int, bool) {
	index := h.hash(key)
	start := index

	for {
		e := h.table[index]
		if e == nil { // empty slot ends search
			return 0, false
		}
		if e.used && !e.deleted && e.key == key {
			return e.value, true
		}
		index = (index + 1) % len(h.table)
		if index == start {
			return 0, false
		}
	}
}

// remove key
func (h *HashMap) Remove(key string) {
	index := h.hash(key)
	start := index

	for {
		e := h.table[index]
		if e == nil {
			return // not found
		}
		if e.used && !e.deleted && e.key == key {
			e.deleted = true // mark tombstone
			h.size--
			return
		}
		index = (index + 1) % len(h.table)
		if index == start {
			return
		}
	}
}

// print
func (h *HashMap) Print() {
	for i, e := range h.table {
		if e != nil && e.used && !e.deleted {
			fmt.Printf("Slot %d: %s=%d\n", i, e.key, e.value)
		} else {
			fmt.Printf("Slot %d: empty\n", i)
		}
	}
}

func main() {
	hm := NewHashMap(5)
	hm.Put("apple", 100)
	hm.Put("banana", 200)
	hm.Put("cherry", 300)

	// get values with proper handling of two returns
	if v, ok := hm.Get("apple"); ok {
		fmt.Println("apple:", v)
	}
	if v, ok := hm.Get("banana"); ok {
		fmt.Println("banana:", v)
	}
	if v, ok := hm.Get("cherry"); ok {
		fmt.Println("cherry:", v)
	}

	hm.Remove("banana")
	fmt.Println("after removing banana:")
	hm.Print()

	// trigger resize
	hm.Put("date", 400)
	hm.Put("fig", 500)
	hm.Put("grape", 600)

	fmt.Println("after adding more:")
	hm.Print()
}
