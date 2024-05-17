package behavioral

import "fmt"

// payments

type PayStrategy interface {
	Pay(amount float64)
}

type CreditCard struct {
	name   string
	number string
}

func (cc *CreditCard) Pay(amount float64) {
	fmt.Printf("Payed %.2f using credit card %s\n", amount, cc.name)
}

type PayPal struct {
	email string
}

func (pp *PayPal) Pay(amount float64) {
	fmt.Printf("Payed using paypal %.2f %s\n", amount, pp.email)
}

type Bitcoin struct {
	walletAddress string
}

func (b *Bitcoin) Pay(amount float64) {
	fmt.Printf("Payed with bitcoin wallet: %s ammount: %.2f\n", b.walletAddress, amount)
}

// Context :Payment
type Payment struct {
	strategy PayStrategy
}

func (p *Payment) SetStrategy(strategy PayStrategy) {
	p.strategy = strategy
}

func (p *Payment) Pay(amount float64) {
	p.strategy.Pay(amount)
}
