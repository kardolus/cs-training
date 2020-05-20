package main

import "fmt"

/*
Provide a way to access the elements of an aggregate object sequentially without
exposing its underlying representation.
*/

func main() {
	// normally you would hide internals and have an Add() method
	collection := NewCollection([]int{1, 2, 3, 4, 5, 100})

	PrintElements(&collection)
}

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

func PrintElements(iterator Iterator) {
	for iterator.HasNext() {
		fmt.Printf("%v\n", iterator.Next())
	}
}

type Collection struct {
	data []int
}

func NewCollection(data []int) Collection {
	return Collection{
		data: data,
	}
}

func (c *Collection) Next() interface{} {
	var result = c.data[len(c.data)-1]
	c.data = c.data[:len(c.data)-1]
	return result
}

func (c *Collection) HasNext() bool {
	return len(c.data) > 0
}
