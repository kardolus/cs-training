package main

import "fmt"

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

func (d *DLL) Remove(key int) {
	if d.head.element.key == key {
		d.head = d.head.next

		if d.head != nil {
			d.head.prev = nil
		}
		d.size--
		return
	}

	focus := d.head

	for focus != nil {
		if key == focus.element.key {
			focus.prev.next = focus.next

			if focus.next != nil {
				focus.next.prev = focus.prev
			}

			d.size--
			return
		}
		focus = focus.next
	}
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

func main() {
	var dll DLL

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

	fmt.Println("String()", dll.String())

	element := dll.Get(4)
	element.value = "a new value"
	fmt.Println("String()", dll.String())

	fmt.Println(`---> Remove(4)`)
	dll.Remove(4)
	fmt.Println("Size()", dll.Size())
	fmt.Println("Contains(4)", dll.Contains(4))
	fmt.Println("String()", dll.String())

	fmt.Println(`---> Remove(3)`)
	dll.Remove(3)
	fmt.Println("Size()", dll.Size())
	fmt.Println("Contains(3)", dll.Contains(3))
	fmt.Println("String()", dll.String())

	fmt.Println(`---> Remove(1)`)
	dll.Remove(1)
	fmt.Println("Size()", dll.Size())
	fmt.Println("Contains(1)", dll.Contains(1))
	fmt.Println("String()", dll.String())
}
