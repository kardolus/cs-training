package main

import (
	"fmt"
	"math"
)

// ------- Doubly Linked List below ------
type Element struct {
	key   int
	value interface{}
}

type node struct {
	next    *node
	prev    *node
	element Element
}

type DLL struct {
	head *node
	size int
}

func (d *DLL) Empty() bool {
	return d.size == 0
}

func (d *DLL) Size() int {
	return d.size
}

func (d *DLL) Add(element Element) {
	d.head = &node{element: element, next: d.head}

	if d.head.next != nil {
		d.head.next.prev = d.head
	}

	d.size++
}

func (d *DLL) Contains(key int) bool {
	return d.Get(key) != nil
}

func (d *DLL) Get(key int) *Element {
	focus := d.head

	for focus != nil {
		if key == focus.element.key {
			return &focus.element
		}
		focus = focus.next
	}

	return nil
}

func (d *DLL) Remove(key int) bool {
	focus := d.head

	for focus != nil {
		if key == focus.element.key {
			if focus.prev != nil {
				focus.prev.next = focus.next
			}

			if focus.next != nil {
				focus.next.prev = focus.prev
			}

			if focus.next == nil && focus.prev == nil {
				d.head = nil
			}

			d.size--
			return true
		}
		focus = focus.next
	}

	return false
}

func (d *DLL) String() string {
	var result string

	focus := d.head
	for focus != nil {
		result += fmt.Sprintf("[key: %v, value: %v]-->", focus.element.key, focus.element.value)
		focus = focus.next
	}

	return result
}

// ------- HashTable below ------

type HashTable struct {
	data []*DLL
	size int
}

const tableSize = 10

func NewHashTable() HashTable {
	return HashTable{
		data: make([]*DLL, tableSize),
	}
}

func (e *Element) HashCode() int {
	// Multiplication method: hashcode = Floor(m (kA mod 1))
	multiplier := float64(0.9)
	fraction := math.Mod((float64(e.key) * multiplier), 1)
	hash := math.Floor(tableSize * fraction)

	return int(hash)
}

func (h *HashTable) Put(key int, value interface{}) {
	element := Element{key: key, value: value}
	hash := element.HashCode()

	if h.data[hash] == nil {
		h.data[hash] = new(DLL)
	}

	match := h.data[hash].Get(key)
	if match != nil {
		match.value = value
	} else {
		h.data[hash].Add(element)
		h.size++
	}
}

func (h *HashTable) Empty() bool {
	return h.size == 0
}

func (h *HashTable) Size() int {
	return h.size
}

func (h *HashTable) Get(key int) interface{} {
	tmp := Element{key: key}
	hash := tmp.HashCode()

	if h.data[hash] == nil || h.data[hash].Get(key) == nil {
		return nil
	}

	return h.data[hash].Get(key).value
}

func (h *HashTable) Contains(key int) bool {
	tmp := Element{key: key}
	hash := tmp.HashCode()
	if h.data[hash] == nil {
		return false
	}

	return h.data[hash].Contains(key)
}

func (h *HashTable) Remove(key int) {
	tmp := Element{key: key}
	hash := tmp.HashCode()

	if h.data[hash].Remove(key) {
		h.size--
	}
	if h.data[hash].Empty() {
		h.data[hash] = nil
	}
}

func (h *HashTable) String() string {
	var result = "\n"

	for _, list := range h.data {
		result += fmt.Sprintf("%v\n", list)
	}

	return result
}

// ----- TESTING ------
func main() {
	var dll DLL

	fmt.Println("--- TESTING DOUBLY LINKED LIST ---")
	fmt.Println("Empty()", dll.Empty())
	fmt.Println("Size()", dll.Size())
	fmt.Println("Contains(3)", dll.Contains(3))

	fmt.Println(`---> Add(Element{3, "hi"})`)
	dll.Add(Element{3, "hi"})
	fmt.Println("Empty()", dll.Empty())
	fmt.Println("Size()", dll.Size())
	fmt.Println("Contains(3)", dll.Contains(3))
	fmt.Println("Contains(4)", dll.Contains(4))

	fmt.Println(`---> Add(Element{4, "hi there"})`)
	dll.Add(Element{4, "hi there"})
	fmt.Println("Contains(4)", dll.Contains(4))
	fmt.Printf("Get(4) %v\n", dll.Get(4))
	fmt.Println("Contains(1)", dll.Contains(1))

	fmt.Println(`---> Add(Element{1, "hi there dude"})`)
	dll.Add(Element{1, "hi there dude"})
	fmt.Println("Contains(1)", dll.Contains(1))

	fmt.Println("String()", &dll)

	element := dll.Get(4)
	element.value = "a new value"
	fmt.Println("String()", &dll)

	fmt.Println(`---> Remove(4)`)
	dll.Remove(4)
	fmt.Println("Size()", dll.Size())
	fmt.Println("Contains(4)", dll.Contains(4))
	fmt.Println("String()", &dll)

	fmt.Println(`---> Remove(3)`)
	dll.Remove(3)
	fmt.Println("Size()", dll.Size())
	fmt.Println("Contains(3)", dll.Contains(3))
	fmt.Println("String()", &dll)

	fmt.Println(`---> Remove(1)`)
	dll.Remove(1)
	fmt.Println("Size()", dll.Size())
	fmt.Println("Contains(1)", dll.Contains(1))
	fmt.Println("String()", &dll)

	fmt.Println("--- TESTING HASHTABLE ---")
	anElement := Element{17, "yo dude"}
	fmt.Printf("HashCode %v: %d\n", anElement, anElement.HashCode())
	anElement = Element{11, "yo dude"}
	fmt.Printf("HashCode %v: %d\n", anElement, anElement.HashCode())
	anElement = Element{10, "yo dude"}
	fmt.Printf("HashCode %v: %d\n", anElement, anElement.HashCode())
	anElement = Element{8978723598723, "yo dude"}
	fmt.Printf("HashCode %v: %d\n", anElement, anElement.HashCode())

	hashTable := NewHashTable()
	fmt.Println("Size()", hashTable.Size())
	fmt.Println("Empty()", hashTable.Empty())

	fmt.Println(`---> Put(3, "what's up")`)
	hashTable.Put(3, "what's up")
	fmt.Println("Size()", hashTable.Size())
	fmt.Println("Empty()", hashTable.Empty())
	fmt.Printf("Print HashTable: %v\n", &hashTable)

	fmt.Println(`---> Put(3, "new value")`)
	hashTable.Put(3, "new value")
	fmt.Println("Size()", hashTable.Size())
	fmt.Println("Empty()", hashTable.Empty())
	fmt.Printf("Print HashTable: %v\n", &hashTable)

	fmt.Println(`---> Put(4, "yo")`)
	hashTable.Put(4, "yo")
	fmt.Println("Size()", hashTable.Size())
	fmt.Println("Empty()", hashTable.Empty())
	fmt.Printf("Print HashTable: %v\n", &hashTable)
	fmt.Println("Get(4)", hashTable.Get(4))
	fmt.Println("Contains(17)", hashTable.Contains(17))
	fmt.Println("Contains(4)", hashTable.Contains(4))

	fmt.Println(`---> Remove(4)`)
	hashTable.Remove(4)
	fmt.Println("Size()", hashTable.Size())
	fmt.Println("Empty()", hashTable.Empty())
	fmt.Printf("Print HashTable: %v\n", &hashTable)

	fmt.Println(`---> Remove(3)`)
	hashTable.Remove(3)
	fmt.Println("Size()", hashTable.Size())
	fmt.Println("Empty()", hashTable.Empty())
	fmt.Printf("Print HashTable: %v\n", &hashTable)
}
