package main

import (
  "fmt"
)

func main() {  
  var set = NewSet(2)
  
  fmt.Println("Size()", set.Size())
  fmt.Println("Empty()", set.Empty())
  fmt.Println("Contains(3)", set.Contains(3))
  
  set.Add(3)
  fmt.Println("Size()", set.Size())
  fmt.Println("Empty()", set.Empty())
  fmt.Println("Contains(3)", set.Contains(3))
  fmt.Println("Contains(4)", set.Contains(4))
  
  set.Add(4)
  fmt.Println("Size()", set.Size())
  fmt.Println("Contains(4)", set.Contains(4))
  
  set.Add(3)
  fmt.Println("Size()", set.Size())
  
  set.Add(15)
  fmt.Println("Size()", set.Size())
  
  set.Add(1)
  fmt.Println("Size()", set.Size())
  
  set.Remove(4)
  fmt.Println("Size()", set.Size())
  fmt.Println("Contains(4)", set.Contains(4))
  
  set.Add(4)
  fmt.Println("Size()", set.Size())
  fmt.Println("Contains(4)", set.Contains(4))
}

/*
  Implementation Below
*/

type Set struct {
  data []int
  size int
}

func NewSet(initialSize int) Set {
  return Set{
    data: make([]int, initialSize),
  }
}

func (s *Set) Empty() bool {
  return s.size == 0
}

func (s *Set) Add(value int) {
  s.copy()
  
  if !s.Contains(value) {
    s.data[s.size] = value
    s.size++
  }
}

func (s *Set) Remove(value int) {
  index := s.indexOf(value)
  if index == -1 {
    return 
  }
  s.data[index] = s.data[s.size - 1]
  s.size--
}

func (s *Set) Size() int {
  return s.size
}

func (s *Set) Contains(value int) bool {
  return s.indexOf(value) != -1
}

func (s *Set) copy() {
  if s.size == len(s.data) {
    newData := make([]int, s.size * 2)
    for index := range s.data {
       newData[index] = s.data[index]
    }
    s.data = newData
  }
}

func (s *Set) indexOf(value int) int { // private, order not guaranteed
  for index, d := range s.data {
    if value == d {
      return index
    }
  }
  return -1
}
