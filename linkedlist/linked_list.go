package linkedlist

import (
	"fmt"
	"strings"
)

// List struct which has the information of the node.
type list struct {
	head   *node
	tail   *node
	length int
}

// Create new linked list.
func NewLinkedList() *list {
	return &list{}
}

// Node struct which has the information to the next one.
type node struct {
	value int
	next  *node
}

// Create new instance of node.
func NewNode(value int) *node {
	return &node{
		value: value,
	}
}

// sub-function in Print()
func doPrint(n *node) string {
	var nested string
	if n != nil {
		nested += "{ "
		nested += "value: " + fmt.Sprint(n.value) + ", "
		nested += "next: " + doPrint(n.next)
		nested += " }"
	} else {
		nested += "null"
	}

	return nested
}

// Custom node JSON marshalling.
func (n *node) Print() string {
	if n == nil {
		return "{}"
	}

	return doPrint(n)
}

// inserts node at the end of linked list.
// returns inserted node.
func (l *list) Push(value int) *node {
	var node = NewNode(value)

	// if linkedlist empty
	if l.head == nil {
		l.head = node
		l.tail = node
		l.length += 1

		return node
	} else { // otherwise
		l.tail.next = node
		l.tail = node
		l.length += 1

		return node
	}
}

// removes node at the end of linked list.
// returns popped node.
func (l *list) Pop() *node {
	// if linkedlist empty
	if l.IsEmpty() {
		return nil
	}

	var poppedNode = l.tail
	// if have 1 node
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.length = 0

		return poppedNode
	}

	// if more than 1 node
	var current = l.head
	var penult *node

	for {
		if current.next == l.tail {
			penult = current
			break
		}

		current = current.next
	}

	penult.next = nil
	l.tail = penult
	l.length -= 1

	return poppedNode
}

// inserts node at given position index of linked list
// returns inserted node.
func (l *list) PushAt(value int, index int) *node {
	if index < 0 || index > l.length {
		return nil
	}

	// create new Node
	var insertedNode = NewNode(value)

	var current = l.head
	var prev *node

	// if linkedlist have 1 node
	if index == 0 {
		if l.head == nil {
			l.tail = insertedNode
		}

		insertedNode.next = l.head
		l.head = insertedNode
		l.length += 1

		return insertedNode
	}

	// if linkedlist have more than 1 node
	// iterates over the linked list to find
	// the position to insert
	var i int

	for i < index {
		prev = current
		current = current.next
		i += 1
	}

	// adding node to linked list
	insertedNode.next = current
	prev.next = insertedNode
	l.length += 1

	if i == l.length-1 {
		l.tail = insertedNode
	}

	return insertedNode
}

// returns node with given index.
func (l *list) Get(index int) *node {
	if index < 0 || index > l.length {
		return nil
	}

	if index == 0 {
		return l.head
	}

	// if given index is greater than 0
	var current = l.head
	var i int

	for i < index {
		current = current.next
		i += 1
	}

	return current
}

// returns deleted node with given index.
func (l *list) Delete(index int) *node {
	if index < 0 || index > l.length {
		return nil
	}

	if index == 0 {
		var deletedNode = l.head
		l.head = l.head.next
		l.length -= 1

		return deletedNode
	}

	// if given index is greater than 0
	var current = l.head
	var prev *node
	var i int

	for i < index {
		prev = current
		current = current.next
		i += 1
	}

	var deletedNode = current
	prev.next = current.next
	l.length -= 1

	return deletedNode
}

// Determines whether linked list has nodes.
func (l *list) IsEmpty() bool {
	return l.length == 0
}

func (l *list) Size() int {
	return l.length
}

// Converts linked list to string with "=>" separated.
func (l *list) ToString() string {
	var values []string
	var current = l.head

	for current != nil {
		values = append(values, fmt.Sprint(current.value))
		current = current.next
	}

	return strings.Join(values, " => ")
}
