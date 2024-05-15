package SOLID

/*
* Objects of a superclass shall be replaceable with objects of its
subclasses without affecting the correctness of the program.
This is more relevant in class-based languages but can apply to Go through its use of interfaces.
* If a function accepts an interface type, any implementation of that interface
should be able to be used without causing the function to fail.
*/

type Bird interface {
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	// Sparrow flying implementation
}

type Ostrich struct{}

// Ostrich does not implement Fly, it shouldn't be used wherever Bird is expected
