package main

import (
	"data-structure-in-golang/array"
	"data-structure-in-golang/linkedlist"
	"data-structure-in-golang/queue"
	"data-structure-in-golang/stack"
	"fmt"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func double(a int) int {
	return add(a, a)
}

func main() {
	// call stack example
	fmt.Println(double(5))
	// DATA STRUCTURE - ARRAY

	fmt.Println("DATA STRUCTURE - ARRAY")
	fmt.Println("=================================================")
	arrNames := array.NewArrayOfString()

	fmt.Println(arrNames.GetAll())

	arrNames.Append("A", "B", "C", "D", "E")
	fmt.Println(arrNames.Pop())
	fmt.Println(arrNames.GetAll())

	fmt.Println(arrNames.GetAllReversed())

	fmt.Println(arrNames.GetOne(2))

	err := arrNames.UpdateOne(3, "Z")
	if err != nil {
		panic(err)
	}
	fmt.Println(arrNames.GetAll())

	err = arrNames.DeleteOne(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(arrNames.GetAll())

	// push to start of array
	arrNames.Prepend("1", "2")
	fmt.Println(arrNames.GetAll())

	newArrNames := arrNames.Map(func(value string) string {
		return value + " - 1"
	})
	fmt.Println(newArrNames.GetAll())

	var txt string
	arrNames.ForEach(func(value string) {
		if value == arrNames[len(arrNames)-1] {
			txt += value
		} else {
			txt += value + " => "
		}
	})

	fmt.Println(txt)

	arrDevices := array.NewArrayOfString()
	arrDevices.Prepend("Mouse", "Keyboard", "Spouse", "SkateBoard", "Laptop")

	fmt.Println(arrDevices.GetAll())

	filteredDevices := arrDevices.Filter(func(value string) bool {
		return strings.Contains(strings.ToLower(value), strings.ToLower("ard"))
	})

	fmt.Println(filteredDevices.GetAll())

	findDevice := arrDevices.Find(func(value string) bool {
		return value == "Spouse"
	})

	fmt.Println(findDevice)
	fmt.Println(arrDevices.Includes("Mouse"))
	fmt.Println(arrDevices.ToString())
	fmt.Println(arrDevices.Shift())
	fmt.Println(arrDevices.ToString())
	fmt.Println(arrDevices.Splice(2, 1, "abc", "def"))
	fmt.Println(arrDevices.GetAll())

	fmt.Println("\n\nDATA STRUCTURE - STACK")
	fmt.Println("=================================================")

	// DATA STRUCTURE - STACK
	s := stack.NewStack()
	s.Push(14)
	s.Push(42)
	s.Push(25)
	fmt.Println("Stack Peek :", s.Peek())
	fmt.Println("Stack Size :", s.Size())
	s.Pop()
	fmt.Println("Stack Peek :", s.Peek())
	fmt.Println("Stack Size :", s.Size())
	s.Pop()
	fmt.Println("Stack Peek :", s.Peek())
	fmt.Println("Stack Size :", s.Size())
	s.Pop()
	fmt.Println("Stack Size :", s.Size())

	fmt.Println("\n\nDATA STRUCTURE - QUEUE")
	fmt.Println("=================================================")

	// DATA STRUCTURE - QUEUE
	q := queue.NewQueue()
	q.Enqueue("`Wathing Movies on Netflix`")
	fmt.Println("Enqueued element : `Wathing Movies on Netflix`")
	q.Enqueue("`Coding Golang`")
	fmt.Println("Enqueued element : `Coding Golang`")
	q.Enqueue("`Reading Books`")
	fmt.Println("Enqueued element : `Reading Books`")
	fmt.Println("====")
	fmt.Println("Peek Queue element :", q.Peek())
	fmt.Println("Size Queue :", q.Size())
	fmt.Println("====")
	fmt.Println("Dequeued element :", q.Dequeue())
	fmt.Println("Peek Queue element :", q.Peek())
	fmt.Println("Size Queue :", q.Size())
	fmt.Println("====")
	fmt.Println("Dequeued element :", q.Dequeue())
	fmt.Println("Peek Queue element :", q.Peek())
	fmt.Println("Size Queue :", q.Size())
	fmt.Println("====")
	fmt.Println("Dequeued element :", q.Dequeue())
	fmt.Println("Size Queue :", q.Size())

	fmt.Println("\n\nDATA STRUCTURE - LINKED LIST")
	fmt.Println("=================================================")

	// Create a list of 5 nodes
	list := linkedlist.NewSinglyLinkedList()
	for i := 1; i <= 10; i++ {
		list.Push(i)
	}

	fmt.Println(list.PushAt(11, 9))
	list.Pop()
	// print JSON marshalled node
	fmt.Println(list.Get(5).Print())

	fmt.Println(list.ToString())

	fmt.Println(list.IsEmpty())
	fmt.Println("Size of linked list :", list.Size())
	list.Delete(1)
	// print linked list to string
	fmt.Println(list.ToString())
}
