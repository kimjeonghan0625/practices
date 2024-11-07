package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val T
}

type DoubleNode[T any] struct {
	prev *DoubleNode[T]
	next *DoubleNode[T]
	val T
}

func SingleLinkedList() {
	root := &Node[int]{nil, 10}
	root.next = &Node[int]{nil, 20}
	root.next.next = &Node[int]{nil, 30}
	
	for n := root; n != nil ; n = n.next {
		fmt.Printf("node val: %d\n", n.val)
	}
}

func DoubleLinkedList() {
	root := &DoubleNode[int]{nil, nil, 10}
	root.next = &DoubleNode[int]{root,nil, 20}
	root.next.next = &DoubleNode[int]{root.next, nil, 30}
	tail := root.next.next
	
	for n := tail; n != nil; n = n.prev {
		fmt.Printf("node val: %d\n", n.val)
	}
}
