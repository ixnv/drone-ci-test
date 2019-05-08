package main

import "testing"

func TestList_Append(t *testing.T) {
	list := List{}

	nodeValue := 1

	node := Node{nodeValue, nil}

	list.Append(&node)

	if list.Head.Value != nodeValue {
		t.Fatal("list head does not equal to just appended node")
	}

	if list.Tail.Value != nodeValue {
		t.Fatal("list tail does not equal to just appended node")
	}
}
