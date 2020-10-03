package main

import (
	"fmt"
)

type Node struct {
	val int
	next *Node
}

func main() {
	var head *Node
	for i := 0; i <= 100; i++ {
		head = insert(head, i)
	}
	fmt.Println(search(head, 99))
	head = remove(head, 2)
	going := false
	for going != true {
		fmt.Printf("%d->", (head).val)
		if head.next == nil {
			going = true
		}
		head = head.next
	}
}

func insert(node *Node, val int) *Node {
	if node == nil {
		head := Node{val: val, next: nil}
		return &head
	}
	head := Node{val: val, next: node}
	return &head
}

func search(head *Node, target int) bool {
	if head == nil {
		return false
	}
	current := head
	found := false
	for found != true {
		current = current.next
		if current.val == target {
			found = true
		}
		if current.next == nil {
			return found
		}
	}
	return found
}

func remove(head *Node, val int) *Node {
	current := head
	found := false
	for found != true {
		if current.next.val == val {
			current.next = current.next.next
			found = true
		} else if head.val == val {
			head = head.next
			return head
		}
		current = current.next
	}
	return head
}