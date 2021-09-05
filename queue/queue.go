package queue

type queue struct {
	// stack [1, 2, 3, 4, 5, 6, 7]
	items []string // [4, 3, 2, 1]
}

func NewQueue() *queue {
	return &queue{}
}

// append new elements at the start of an array
func (q *queue) Enqueue(s string) {
	var arr []string
	arr = append(arr, s)
	q.items = append(arr, q.items...)
}

// return dequeued string
func (q *queue) Dequeue() string {
	item, items := q.items[len(q.items)-1], q.items[:len(q.items)-1]
	q.items = items
	return item
}

// get last element of queue
func (q *queue) Peek() string {
	return q.items[len(q.items)-1]
}

// return size of queue
func (q *queue) Size() int {
	return len(q.items)
}
