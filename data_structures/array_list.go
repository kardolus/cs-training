package main

import "fmt"

type List struct {
	data []int
	size int
}

const initialSize = 1

func NewList() List {
	var result List

	result.data = make([]int, initialSize)

	return result
}

func (l *List) Empty() bool { // O(1)
	return l.size == 0
}

func (l *List) Add(value int) { // O(1)
	l.copy()
	l.data[l.size] = value
	l.size++
}

func (l *List) Insert(index, value int) { // O(N)
	l.copy()

	for i := index; i < l.size-1; i++ {
		l.data[i+1] = l.data[i] // shift forward
	}

	l.data[index] = value
	l.size++
}

func (l *List) Remove(value int) { // O(N)
	var match bool
	for index := 0; index < l.size; index++ {
		if l.data[index] == value {
			l.size--
			match = true
		}
		if match && index < len(l.data)-1 {
			l.data[index] = l.data[index+1] // shift backward
		}
	}
}

func (l *List) Contains(value int) bool { // O(N)
	for i := 0; i < l.size; i++ {
		if l.data[i] == value {
			return true
		}
	}

	return false
}

func (l *List) Size() int { // O(1)
	return l.size
}

func (l *List) copy() {
	if len(l.data) == l.size {
		newData := make([]int, l.size*2)
		copy(newData, l.data)
		l.data = newData
	}
}

// Run tests
func main() {
	var list = NewList()
	fmt.Println("Empty?", list.Empty())

	list.Add(3)
	fmt.Println("Empty?", list.Empty())
	fmt.Println("Contains(3)?", list.Contains(3))
	fmt.Println("Contains(4)?", list.Contains(4))

	fmt.Println("Size == 1?", list.Size() == 1)
	list.Add(4)
	fmt.Println("Contains(4)?", list.Contains(4))
	fmt.Println("Size == 2?", list.Size() == 2)

	list.Remove(3)
	fmt.Println("Size == 1?", list.Size() == 1)
	fmt.Println("Contains(3)?", list.Contains(3))

	list.Add(3)
	fmt.Println("Contains(3)?", list.Contains(3))

	list.Remove(3)
	fmt.Println("Size == 1?", list.Size() == 1)
	fmt.Println("Contains(3)?", list.Contains(3))

	list.Add(7)
	list.Add(8)
	list.Add(9)
	fmt.Println(list.data, list.size)
	list.Remove(7)
	fmt.Println(list.data, list.size)

	list.Insert(1, 5)
	fmt.Println(list.data, list.size)
}
