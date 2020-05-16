package main

import (
	"fmt"
	"strconv"
)

func main() {
	var list LinkedList
	fmt.Println("Empty?", list.Empty())

	list.Add(3)
	fmt.Println(list.PrintList())
	fmt.Println("Empty?", list.Empty())
	fmt.Println("Size", list.Size())
	fmt.Println("Contains(3)?", list.Contains(3))
	fmt.Println("Contains(4)?", list.Contains(4))

	list.Add(4)
	fmt.Println(list.PrintList())
	fmt.Println("Empty?", list.Empty())
	fmt.Println("Size", list.Size())
	fmt.Println("Contains(3)?", list.Contains(3))
	fmt.Println("Contains(4)?", list.Contains(4))

	list.Add(2)
	list.Add(5)
	list.Add(8)
	list.Add(4)
	fmt.Println(list.PrintList())
	fmt.Println("Size", list.Size())
	fmt.Println("Contains(8)?", list.Contains(8))
	list.Remove(8)
	fmt.Println(list.PrintList())
	fmt.Println("Size", list.Size())
	fmt.Println("Contains(8)?", list.Contains(8))

	list.Remove(4)
	fmt.Println(list.PrintList())
	fmt.Println("Size", list.Size())
	fmt.Println("Contains(4)?", list.Contains(4))
}

/*
* Implementation starts here
 */

type LinkedList struct {
	head *node
	size int
}

type node struct {
	value int
	next  *node // to get around "invalid recursive type node"
}

func (l *LinkedList) Add(value int) { // O(1)
	newHead := &node{
		value: value,
		next:  l.head,
	}

	l.head = newHead
	l.size++
}

func (l *LinkedList) Size() int { // O(1)
	return l.size
}

func (l *LinkedList) Empty() bool { // O(1)
	return l.head == nil
}

func (l *LinkedList) Remove(value int) { // O(n)
	// node(3)-node(4)-node(2)-node(5)
	// Delete: 2
	// node(4).next = node(5)

	var (
		prev  *node
		focus = l.head
	)

	for focus != nil {
		if l.head.value == value {
			l.head = l.head.next
			l.size--
		}

		if focus.value == value && prev != nil {
			prev.next = focus.next
			l.size--
		}

		prev = focus
		focus = focus.next
	}
}

func (l *LinkedList) Contains(value int) bool { // O(n)
	focus := l.head

	for focus != nil {
		if focus.value == value {
			return true
		}
		focus = focus.next
	}

	return false
}

func (l *LinkedList) PrintList() string {
	var result string

	focus := l.head

	for focus != nil {
		result = strconv.Itoa(focus.value) + "-" + result
		focus = focus.next
	}

	return result
}
