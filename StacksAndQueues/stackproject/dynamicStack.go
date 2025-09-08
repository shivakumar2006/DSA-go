package main

// dynamic stack embeds CustomStack to inherit it.
type DynamicStack struct {
	*CustomStack
}

func NewDynamicStack(capacity int) *DynamicStack {
	return &DynamicStack{CustomStack: NewCustomStack(capacity)}
}

// Override push to make it dynamic
func (d *DynamicStack) Push(value int) error {
	if len(d.data) >= d.size {
		d.size *= 2 // double the capacity
	}
	return d.CustomStack.Push(value) // otherwise, call the original path
}
