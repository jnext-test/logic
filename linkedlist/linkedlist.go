package linkedlist

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (ll *List) Add(data int) {
	newNode := &Node{
		data: data,
	}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (ll *List) Display() {
	current := ll.head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func (ll *List) Delete(data int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (ll *List) DeleteLast() {
	if ll.head == nil { // Case 1: Empty list
		return
	}

	if ll.head.next == nil { // Case 2: Single node list
		ll.head = nil
		return
	}

	// Case 3: Multiple nodes
	current := ll.head
	for current.next.next != nil { // Traverse to the second-to-last node
		current = current.next
	}
	current.next = nil // Set the next of the second-to-last node to nil
}
