package main

import "fmt"

// Observer interface declares the method to update observers -->for loose-coupling any type implement
// this can use update
type Observer interface {
	Update(newState interface{})
}

// Subject maintains a list of observers and notifies them of state changes
type Subject struct {
	observers []Observer
	state     interface{}
}

// RegisterObserver adds a new observer to the subject's list
func (s *Subject) RegisterObserver(o Observer) {
	s.observers = append(s.observers, o)
}

// RemoveObserver removes an observer from the subject's list
func (s *Subject) RemoveObserver(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers informs all observers about state changes
func (s *Subject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

// ConcreteObserver implements the Observer interface and reacts to state changes
type ConcreteObserver struct {
	id int // an identifier for the observer, for instance

}

// Update reacts to state changes in the subject
func (co *ConcreteObserver) Update(newState interface{}) {
	fmt.Println("----------------------")
	fmt.Println("New state sent:", newState)
	fmt.Println("Observer notified--> Observer id:", co.id)
	fmt.Println("----------------------")
}

func main() {
	// Instantiate a subject
	subject := &Subject{}

	// Instantiate and register observers
	observer1 := &ConcreteObserver{id: 1}
	observer2 := &ConcreteObserver{id: 2}

	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	// Change the subject's state and notify observers
	newState := "State-changed"
	subject.state = newState
	subject.NotifyObservers()
}
