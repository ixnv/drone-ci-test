package main

import "fmt"

type (
	Node struct {
		Value interface{} `json:"value"`
		Next  *Node       `json:"next"`
	}

	List struct {
		Head *Node `json:"head"`
		Tail *Node `json:"tail"`
		}
)

func (l *List) Append(newNode *Node) {
	if l.Head == nil {
		l.Head = newNode
		l.Tail = l.Head
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func (l *List) print() {
	for it := l.Head; it != nil; it = it.Next {
		fmt.Println(it.Value)
	}
}
