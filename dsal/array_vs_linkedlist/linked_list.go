package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

type DoubleNode[T any] struct {
	prev *DoubleNode[T]
	next *DoubleNode[T]
	val  T
}

func SingleLinkedList() {
	root := &Node[int]{nil, 10}
	root.next = &Node[int]{nil, 20}
	root.next.next = &Node[int]{nil, 30}

	for n := root; n != nil; n = n.next {
		fmt.Printf("node val: %d\n", n.val)
	}
}

func DoubleLinkedList() {
	root := &DoubleNode[int]{nil, nil, 10}
	root.next = &DoubleNode[int]{root, nil, 20}
	root.next.next = &DoubleNode[int]{root.next, nil, 30}
	tail := root.next.next

	for n := tail; n != nil; n = n.prev {
		fmt.Printf("node val: %d\n", n.val)
	}
}

func AddNode[T any](root *Node[T], next *Node[T]) *Node[T] {
	root.next = next
	return next
}

func PrintNode[T any](start *Node[T]) {
	root := start
	for root != nil {
		fmt.Printf("Value is : %v\n", root.val)
		root = root.next
	}
}

func (r *Node[int]) InsertNode(dist uint, item *Node[int]) *Node[int] {
	if dist == 0 {
		item.next = r
		return item
	}
	tmp := r
	var prev *Node[int]
	for i := 0; uint(i) < dist; i++ {
		if dist-1 == uint(i) {
			prev = tmp
		}
		if tmp == nil {
			fmt.Println("리스트의 범위를 초과한 삽입 요청입니다.")
			return item
		}
		tmp = tmp.next
	}
	prev.next = item
	item.next = tmp
	return item
}

func LLInsertDemo() {
	tmproot := &Node[int]{nil, 10}
	root := tmproot
	for i := 2; i < 5; i++ {
		tmproot = AddNode(tmproot, &Node[int]{nil, i * 10})
	}
	PrintNode(root)

	root.InsertNode(4, &Node[int]{nil, 100})

	fmt.Println()
	fmt.Println("Insert Complete.")
	PrintNode(root)
}
