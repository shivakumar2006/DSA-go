// Open Adressing in hashmaps

package main

type entry struct {
	key   string
	value int
	used  bool // to mark that this slot is occupied
}

type hashmap struct {
	table []*entry
	size  int
}

func NewHashMap(capacity int) *hashmap {
	return &hashmap{
		table: make([]*entry, capacity),
	}
}

func (h *hashmap) Hash(key string) int {
	hash := 0
	for _, c := range key {
		hash += int(c)
	}
	return hash % len(h.table)
}
