package main

import (
	"fmt"
	"hash/fnv"
)

type Entity struct {
	key   interface{}
	value interface{}
	next  *Entity
}

// hashmapfinal is our custom hashmap
type HashmapFinal struct {
	buckets    []*Entity
	size       int
	loadFactor float64
}

func NewHashMapFinal() *HashmapFinal {
	buckets := make([]*Entity, 10)
	return &HashmapFinal{
		buckets:    buckets,
		size:       0,
		loadFactor: 0.5,
	}
}

// hash function for keys
func (hm *HashmapFinal) hash(key interface{}) int {
	h := fnv.New64a()
	fmt.Fprintf(h, "%v", key)
	return int(h.Sum64() % uint64(len(hm.buckets)))
}

// put inserts or update a new value pair
func (hm *HashmapFinal) Put(key, value interface{}) {
	if float64(hm.size)/float64(len(hm.buckets)) > hm.loadFactor {
		hm.rehash()
	}

	index := hm.hash(key)
	head := hm.buckets[index]

	// check if key exists
	for e := head; e != nil; e = e.next {
		if e.key == key {
			e.value = value
			return
		}
	}

	// insert new entity at head of the list
	newEntity := &Entity{key: key, value: value, next: head}
	hm.buckets[index] = newEntity
	hm.size++
}

// Get retrieves the values by key
func (hm *HashmapFinal) Get(key interface{}) (interface{}, bool) {
	index := hm.hash(key)
	for e := hm.buckets[index]; e != nil; e = e.next {
		if e.key == key {
			return e.value, true
		}
	}
	return nil, false
}

// remove deletes akey valeue pair
func (hm *HashmapFinal) Remove(key interface{}) {
	index := hm.hash(key)
	head := hm.buckets[index]
	var prev *Entity

	for e := head; e != nil; e = e.next {
		if e.key == key {
			if prev == nil {
				hm.buckets[index] = e.next
			} else {
				prev.next = e.next
			}
			hm.size--
			return
		}
		prev = e
	}
}

// contains key if a key exists
func (hm *HashmapFinal) ContainsKey(key interface{}) bool {
	_, ok := hm.Get(key)
	return ok
}

// rehash, double the bucket array and reinsets reshash
func (hm *HashmapFinal) rehash() {
	fmt.Println("We are now rehashing!")

	oldBuckets := hm.buckets
	hm.buckets = make([]*Entity, len(oldBuckets)*2)
	hm.size = 0

	for _, head := range oldBuckets {
		for e := head; e != nil; e = e.next {
			hm.Put(e.key, e.value)
		}
	}
}

// string returns string representation
func (hm *HashmapFinal) String() string {
	result := "{"
	for _, head := range hm.buckets {
		for e := head; e != nil; e = e.next {
			result += fmt.Sprintf("%v = %v, ", e.key, e.value)
		}
	}
	result += "}"
	return result
}

// iterate prints all key value pairs
func (hm *HashmapFinal) Iterate() {
	for i, head := range hm.buckets {
		fmt.Printf("Bucket %d: ", i)
		for e := head; e != nil; e = e.next {
			fmt.Printf("[%v = %v] -> ", e.key, e.value)
		}
		fmt.Println("nil")
	}
}

func main() {
	hm := NewHashMapFinal()

	hm.Put("Shiva", 19)
	hm.Put("Vishal", 17)
	hm.Put("Priyam", 22)

	fmt.Println("Map : ", hm)
	hm.Iterate()

	val, ok := hm.Get("Shiva")
	if ok {
		fmt.Println("Shiva's age : ", val)
	}

	hm.Put("Shiva", 20) // update
	fmt.Println("After updating shiva : ", hm)
	hm.Iterate()

	fmt.Println("contains Vishal?", hm.ContainsKey("Vishal"))

	hm.Remove("Priyam")
	fmt.Println("After removing priyam : ", hm)
	hm.Iterate()

	// Add more to trigger rehash
	hm.Put("A", 1)
	hm.Put("B", 2)
	hm.Put("C", 3)
	hm.Put("D", 4)
	hm.Put("E", 5)
	hm.Put("F", 6)

	fmt.Println("After adding more elements (trigger rehash if needed) : ", hm)
	hm.Iterate()
}
