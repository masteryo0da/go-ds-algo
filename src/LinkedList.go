package src

import "fmt"

type Node struct {
	element int
	next *Node
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) Add(element int) {
	if l.root == nil {
		l.root = &Node{element: element}
		l.tail = l.root
	} else {
		l.tail.next = &Node{element: element}
		l.tail = l.tail.next
	}
}

func (l *LinkedList) Print() {
	temp := l.root
	for temp != nil {
		fmt.Println(temp.element)
		temp = temp.next
	}
}
