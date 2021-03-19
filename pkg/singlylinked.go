package main

import (
	"fmt"
)

type Node struct {
	key int
	next *Node
}

type Head struct {
	node *Node
}

func Add(key int, head *Head) {
	node := new(Node)
	node.key = key
	if ( head.node == nil ) {
		head.node = node
	} else {
		temp := head.node
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = node
	}
	
}

func PrintList(head *Head) bool{
	temp := head.node
	if temp == nil {
		fmt.Println("Linked list Empty")
		return false
	}
	for temp != nil {
		fmt.Println(temp.key)
		temp = temp.next
	}
	return true
}

func main() {
	head := new(Head)
	Add(1, head)
	Add(2, head)
	Add(3, head)
	Add(4, head)
	PrintList(head)	
}
