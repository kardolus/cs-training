package main

import (
	"fmt"
)

/*
	Like the View in MVC. Lets Objects subscribe to changes in a Subject.
*/

func main() {
	var subject Subject

	NewObserver(&subject)

	subject.SetState("crushing")
}

type Observer interface {
	Notify()
}

type Subject struct {
	observers []Observer
	state     string // for simplicity
}

func (s *Subject) SetState(state string) {
	s.state = state

	for _, observer := range s.observers {
		observer.Notify()
	}
}

func (s *Subject) GetState() string {
	return s.state
}

func (s *Subject) Subscribe(observer Observer) {
	s.observers = append(s.observers, observer)
}

type SomeObserver struct {
	subject *Subject
}

func NewObserver(subject *Subject) {
	var result = SomeObserver{
		subject: subject,
	}

	subject.Subscribe(result)
}

func (s SomeObserver) Notify() {
	fmt.Println("Thanks! Let me get the state")
	fmt.Println("State:", s.subject.GetState())
}
