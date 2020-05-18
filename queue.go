package main

import (
  "fmt"
)

func main() {  
  var queue = NewQueue(2)
  
  fmt.Println("Empty()", queue.Empty())
  fmt.Println("Size()", queue.Size())
  fmt.Println("Contains(3)", queue.Contains(3))
  fmt.Println("Peek() %+v", wrap(queue.Peek()))
  fmt.Println("Poll() %+v", wrap(queue.Poll()))
  
  fmt.Println("----> Offer(3)")
  queue.Offer(3)
  
  fmt.Println("Empty()", queue.Empty())
  fmt.Println("Size()", queue.Size())
  fmt.Println("Contains(3)", queue.Contains(3))
  fmt.Println("Peek() %+v", wrap(queue.Peek()))
  fmt.Println("Poll() %+v", wrap(queue.Poll()))
  fmt.Println("Size()", queue.Size())
  
  fmt.Println("----> Offer(3)")
  queue.Offer(3)
  
  fmt.Println("----> Offer(4)")
  queue.Offer(4)
  
  fmt.Println("Contains(4)", queue.Contains(4))
  fmt.Println("Peek() %+v", wrap(queue.Peek()))
  fmt.Println("Size()", queue.Size())
  fmt.Println("Poll() %+v", wrap(queue.Poll()))
  fmt.Println("Size()", queue.Size())
  fmt.Println("Contains(4)", queue.Contains(4))
  fmt.Println("Contains(3)", queue.Contains(3))
}

func wrap(vs ...interface{}) []interface{} {
    return vs
}

/*
  Implementation below
*/

type Queue struct {
  data []int
  tail int
  head int
}

func NewQueue(capacity int) Queue {
  return Queue{
    data: make([]int, capacity),
  }
}

func (q *Queue) Empty() bool {
  return q.tail == 0
}

func (q *Queue) Size() int {
  return q.tail - q.head
}

func (q *Queue) Offer(value int) {
  q.copy()
  q.data[q.tail] = value
  q.tail++ 
}

func (q *Queue) Contains(value int) bool {
  for i := q.head; i < q.tail; i++ {
    if value == q.data[i] {
      return true
    }
  }
  return false
}

func (q *Queue) Peek() (int, error) {
  if q.tail - q.head == 0 {
    return 0, fmt.Errorf("the queue is empty")
  }
  return q.data[q.head], nil
}

func (q *Queue) Poll() (int, error) {
  result, err := q.Peek()
  if err != nil {
    return 0, err
  }
  
  q.head++
  
  return result, nil
}

func (q *Queue) copy() {
  if q.tail == len(q.data) {
    newData := make([]int, q.tail*2)
    copy(newData, q.data)
    q.data = newData
  }
}
