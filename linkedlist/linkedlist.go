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
