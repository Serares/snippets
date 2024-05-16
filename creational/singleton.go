package creational

// Create a singleton pattern in golang
type Singleton struct {
	value int
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{value: 0}
	}
	return instance
}
