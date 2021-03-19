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

func Delete(key int, head *Head) bool{
	temp := head.node
	if temp == nil {
		fmt.Println("List empty cant delete")
		return false
	} else if temp.key == key {
		fmt.Println("Key Found")
		head.node = temp.next
	}else {
		for temp.next != nil {
			prev := temp
			if temp.next.key == key {
				fmt.Println("Key Found")
				prev.next = temp.next.next
				return true
				
			}
			temp = temp.next
		}
	}
	return false
	
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
	Delete(1, head)
	PrintList(head)
	
}
