package main

import (
	"fmt"
	"sync"
)

var (
	mutex     = &sync.Mutex{}
	expensive *ExpensiveObject
)

func main() {
	var proxy *ProxyObject
	DoingSomethingExpensive(proxy, 3)
	DoingSomethingExpensive(proxy, 4)
}

type IExpensive interface {
	DoSomething(value int)
}

type ExpensiveObject struct {
}

func (e *ExpensiveObject) DoSomething(value int) {
	fmt.Println("Doing it")
}

type ProxyObject struct {
}

func (e *ProxyObject) DoSomething(value int) {
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("Adding some logging here: ", value)

	if value == 4 { // Because for some reason we only care about 4
		if expensive == nil { // So we have just a singleton of the expensive object
			expensive = new(ExpensiveObject)
		}
		expensive.DoSomething(value)
	}
}

func DoingSomethingExpensive(e IExpensive, value int) {
	e.DoSomething(value)
}
