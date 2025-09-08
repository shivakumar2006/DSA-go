package main

// dynamic queue embeded form custom queue
type DynamicQueue struct {
	*CustomQueue
}

func NewdynamicQueue(capacity int) *DynamicQueue {
	return &DynamicQueue{CustomQueue: NewCustomQueue(capacity)}
}

// override push to make it dynamic
func (q *DynamicQueue) Push(value int) error {
	if len(q.data) >= q.size {
		// double the size
		q.size *= 2
	}
	// otherwise, call original enqueue
	return q.CustomQueue.Enqueue(value)
}
