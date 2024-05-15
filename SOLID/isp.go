package SOLID

// Interface Segregation Principle
/*
* Clients should not be forced to depend upon interfaces they do not use.
This means having several specific interfaces is better than one general-purpose interface.
* Instead of one large interface, you should define multiple smaller interfaces.
*/
type Walker interface {
	Walk()
}

type Swimmer interface {
	Swim()
}

type Duck struct{}

func (d Duck) Walk() {
	// Duck walking
}

func (d Duck) Swim() {
	// Duck swimming
}
