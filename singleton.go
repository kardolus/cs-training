package main

import (
  "fmt"
  "sync"
)

func main() {  
  fmt.Printf("instance: %+v\n", GetInstance(6))
  fmt.Printf("instance: %+v\n", GetInstance(3))
  fmt.Printf("instance: %+v\n", GetInstance(15))
}

/*
  Implementation below
*/

var (
  single *Singleton
  mutex = &sync.Mutex{}
)

type Singleton struct {
  value int
}

func GetInstance(value int) Singleton {
  mutex.Lock()
  defer mutex.Unlock()
  
  var result Singleton
  
  if single == nil { 
    result = Singleton{value: value}
    single = &result
  } else {
    result = *single
  }
  
  return result
}
