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

func TestStrategy(t *testing.T) {

	p := &Payment{}

	cc := &CreditCard{name: "Jon Do", number: "1234-123123-123123-12313"}
	pp := &PayPal{
		email: "jon.do@domail.com",
	}
	btc := &Bitcoin{
		walletAddress: "12312dwdad12dawse4312",
	}

	p.SetStrategy(cc)

	p.Pay(3.21)

	p.SetStrategy(pp)
	p.Pay(9.2)

	p.SetStrategy(btc)
	p.Pay(1000.233)
}
