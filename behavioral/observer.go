package behavioral

import "fmt"

/*
The Observer pattern defines a one-to-many dependency between

	objects so that when one object changes state,
	all its dependents are notified and updated automatically.
*/
type Observer interface {
	Update(string) string
}

type Subject struct {
	observers []Observer
	state     string
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) NotifyAll() {
	for _, obs := range s.observers {
		obs.Update(s.state)
	}
}

func (s *Subject) SetState(state string) {
	s.state = state
	s.NotifyAll()
}

type ServiceCustomer struct {
	Id int
}

func (sc *ServiceCustomer) Update(state string) string {
	fmt.Printf("Updated the customer service with id %d - state %s\n", sc.Id, state)
	return fmt.Sprintf("Update customer service with id %d", sc.Id)
}

type ServiceOrder struct {
	Id int
}

func (so *ServiceOrder) Update(state string) string {
	fmt.Printf("Update the order service with id %d - state %s\n", so.Id, state)
	return fmt.Sprintf("Update service order with id %d", so.Id)
}
