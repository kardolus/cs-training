package main

import (
	"fmt"
)

/*
* Undirected graph example:
* [0]----[1]
*  |    / | \
*  |   /  |  \
*  |  /   |  [2]
*  | /    |  /
*  |/     | /
* [4]----[3]
 */

/*
 * Adjecency-list representation:
 * [0] --> 1 --> 4
 * [1] --> 0 --> 2 --> 3 --> 4
 * [2] --> 1 --> 3
 * [3] --> 1 --> 2 --> 4
 * [4] --> 0 --> 1 --> 3
 */

type SimpleList struct {
	head    *Node
	visited bool
}

type Node struct {
	next  *Node
	value int
}

func (s *SimpleList) Add(value int) {
	s.head = &Node{
		value: value,
		next:  s.head,
	}
}

func (s *SimpleList) String() string {
	var result string

	focus := s.head

	for focus != nil {
		result += fmt.Sprintf("[%d]-->", focus.value)
		focus = focus.next
	}

	result += "[nil]"

	return result
}

func DepthFirstSearch(adjecencyList []*SimpleList, startIndex int) {
	if adjecencyList[startIndex].visited == true {
		return
	}

	focus := adjecencyList[startIndex].head

	for {
		fmt.Printf("[%d] ", startIndex)
		visit(focus)
		adjecencyList[focus.value].visited = true
		focus = focus.next
		if focus == nil {
			break
		}

		DepthFirstSearch(adjecencyList, focus.value)
	}
}

func visit(node *Node) {
	fmt.Println(node.value)
}

// TODO breadth first search

func main() {
	var list SimpleList
	list.Add(1)
	list.Add(4)

	fmt.Println("Testing the Simple List:", &list)

	adjecencyList := make([]*SimpleList, 5)
	adjecencyList[0] = new(SimpleList)
	adjecencyList[0].Add(1)
	adjecencyList[0].Add(4)
	adjecencyList[0].Add(0)

	adjecencyList[1] = new(SimpleList)
	adjecencyList[1].Add(0)
	adjecencyList[1].Add(2)
	adjecencyList[1].Add(3)
	adjecencyList[1].Add(4)
	adjecencyList[1].Add(1)

	adjecencyList[2] = new(SimpleList)
	adjecencyList[2].Add(1)
	adjecencyList[2].Add(3)
	adjecencyList[2].Add(2)

	adjecencyList[3] = new(SimpleList)
	adjecencyList[3].Add(1)
	adjecencyList[3].Add(2)
	adjecencyList[3].Add(4)
	adjecencyList[3].Add(3)

	adjecencyList[4] = new(SimpleList)
	adjecencyList[4].Add(0)
	adjecencyList[4].Add(1)
	adjecencyList[4].Add(3)
	adjecencyList[4].Add(4)

	fmt.Println("\n------ Adjecency List ------")
	for i := range adjecencyList {
		if adjecencyList[i] != nil {
			fmt.Printf("%s\n", adjecencyList[i])
		}
	}

	fmt.Println("\n------ Depth First Search ------")
	DepthFirstSearch(adjecencyList, 0)
}
