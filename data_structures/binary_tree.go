package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

type BinaryTree struct {
	head *node
}

func (b *BinaryTree) Add(value int) {
	focus := b.head
	if focus == nil {
		b.head = &node{value: value}
		return
	}

	for {
		if focus.value > value {
			if focus.left == nil {
				focus.left = &node{value: value}
				return
			}
			focus = focus.left
		} else {
			if focus.right == nil {
				focus.right = &node{value: value}
				return
			}
			focus = focus.right
		}
	}
}

// The minimum of the right child or the first parent on the right (larger value parent)
func (b *BinaryTree) Successor(value int) int {
	var parents []*node
	focus := b.head

	for focus != nil {
		if focus.value == value {
			if focus.right != nil {
				return Minimum(focus.right)
			}
			break
		}

		parents = append(parents, focus)

		if focus.value > value {
			focus = focus.left
		} else {
			focus = focus.right
		}
	}

	for i := len(parents) - 1; i >= 0; i-- {
		if parents[i].value >= value {
			return parents[i].value
		}
	}

	return 0 // in the real world this would return an error
}

// Maximum of left child or the first left parent (smaller value)
func (b *BinaryTree) Predecessor(value int) int {
	var parents []*node
	focus := b.head

	for focus != nil {
		if focus.value == value {
			if focus.left != nil {
				return Maximum(focus.left)
			}
			break
		}

		parents = append(parents, focus)

		if focus.value > value {
			focus = focus.left
		} else {
			focus = focus.right
		}
	}

	for i := len(parents) - 1; i >= 0; i-- {
		if parents[i].value <= value {
			return parents[i].value
		}
	}

	return 0
}

func (b *BinaryTree) Contains(value int) bool {
	focus := b.head

	for focus != nil {
		if focus.value == value {
			return true
		}
		if focus.value > value {
			focus = focus.left
			continue
		}
		focus = focus.right
	}

	return false
}

func ContainsRecursive(focus *node, value int) bool {
	if focus == nil {
		return false
	}
	if focus.value == value {
		return true
	}

	if focus.value > value {
		return ContainsRecursive(focus.left, value)
	}

	return ContainsRecursive(focus.right, value)
}

func InOrderWalk(focus *node) string {
	var result string

	if focus == nil {
		return result
	}

	result += InOrderWalk(focus.left)
	result += fmt.Sprintf("[%d] ", focus.value)
	result += InOrderWalk(focus.right)

	return result
}

func Minimum(focus *node) int {
	var result int

	for focus != nil {
		result = focus.value
		focus = focus.left
	}

	return result
}

func Maximum(focus *node) int {
	var result int

	for focus != nil {
		result = focus.value
		focus = focus.right
	}

	return result
}

func BreadthFrist(focus *node) {
	queue := []*node{focus}

	i := 0

	for i < len(queue) {
		if queue[i].left != nil {
			queue = append(queue, queue[i].left)
		}
		if queue[i].right != nil {
			queue = append(queue, queue[i].right)
		}
		i++
	}

	for _, item := range queue {
		fmt.Printf("[%d]", item.value)
	}
	fmt.Println()
}

func main() {
	var tree BinaryTree
	var input = []int{20, 8, 25, 34, 9, 5, 50, 40, 7, 15, 8, 51, 47}

	fmt.Println("Adding:", input)
	for _, value := range input {
		tree.Add(value)
	}

	fmt.Println("sorted:", InOrderWalk(tree.head))
	fmt.Println("Minimum()", Minimum(tree.head))
	fmt.Println("Maximum()", Maximum(tree.head))
	fmt.Println("Contains(15)", tree.Contains(15))
	fmt.Println("Contains(16)", tree.Contains(16))
	fmt.Println("Contains(51)", tree.Contains(51))
	fmt.Println("Contains(7)", tree.Contains(7))
	fmt.Println("ContainsRecursive(7)", ContainsRecursive(tree.head, 7))
	fmt.Println("ContainsRecursive(16)", ContainsRecursive(tree.head, 16))
	fmt.Println("ContainsRecursive(51)", ContainsRecursive(tree.head, 51))
	fmt.Println("ContainsRecursive(52)", ContainsRecursive(tree.head, 52))
	fmt.Println("Successor(15)", tree.Successor(15))
	fmt.Println("Successor(8)", tree.Successor(8))
	fmt.Println("Successor(9)", tree.Successor(9))
	fmt.Println("Predecessor(9)", tree.Predecessor(9))
	fmt.Println("Predecessor(15)", tree.Predecessor(15))
	fmt.Println("Predecessor(50)", tree.Predecessor(50))
	fmt.Println("BreadthFirst: --->")
	BreadthFrist(tree.head)
}
