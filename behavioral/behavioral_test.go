package behavioral

import "testing"

func TestObserver(t *testing.T) {

	sc1 := ServiceCustomer{Id: 1}
	so1 := ServiceOrder{Id: 1}

	s1 := Subject{}

	s1.Attach(&sc1)
	s1.Attach(&so1)

	t.Run("Testing the Observer", func(t *testing.T) {
		s1.SetState("Notification")
	})
}
