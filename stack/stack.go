package stack

type stack struct {
	items []int
}

// create new instance of stack.
func NewStack() *stack {
	return &stack{}
}

// push an element to stack.
func (s *stack) Push(item int) {
	s.items = append(s.items, item)
}

// return popped element.
func (s *stack) Pop() int {
	item, items := s.items[len(s.items)-1], s.items[0:len(s.items)-1]
	s.items = items
	return item
}

// get last element of stack.
func (s *stack) Peek() int {
	return s.items[len(s.items)-1]
}

// return size of stack.
func (s *stack) Size() int {
	return len(s.items)
}
