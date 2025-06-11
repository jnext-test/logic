package main

import (
	"logic/linkedlist"
)

func main() {
	ll := &linkedlist.List{}
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	ll.Display()

}
