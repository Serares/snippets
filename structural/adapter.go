package structural

// create a adapter pattern example in golang

type Target interface {
	Request() string
}

type Adaptee struct{}

func (s *Adaptee) SpecificRequest() string {
	return "Specific Request"
}

type Adapter struct {
	Adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}
