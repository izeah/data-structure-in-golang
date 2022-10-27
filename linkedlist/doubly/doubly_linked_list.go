package doubly

import (
	"fmt"
	"strings"
)

// Doubly Linked List struct which has the information of the node.
type list struct {
	head   *node
	tail   *node
	length int
}

// Node struct which has the information to the next one and the previous one.
type node struct {
	value int
	next  *node
	prev  *node
}

// Create new instance of doubly linked list.
func New() *list {
	return &list{}
}

// Create new instance of node with previous.
func newNode(value int) *node {
	return &node{
		value: value,
	}
}

// inserts node at the end of doubly linked list.
// returns inserted node.
func (l *list) Push(value int) *node {
	var insertedNode = newNode(value)

	// if linkedlist empty
	if l.head == nil {
		l.head = insertedNode
		l.tail = insertedNode
		l.length += 1

		return insertedNode
	} else { // otherwise
		var current = l.head
		var prev *node

		for current != nil {
			prev = current
			current = current.next
		}

		insertedNode.prev = prev
		l.tail.next = insertedNode
		l.tail = insertedNode
		l.length += 1

		return insertedNode
	}
}

// Converts doubly linked list to string with "<=>" separated.
func (l *list) String() string {
	var values []string
	var current = l.head

	for current != nil {
		values = append(values, fmt.Sprint(current.value))
		current = current.next
	}

	return strings.Join(values, " <=> ")
}
