package SOLID

// Open Closed Principle
/*
* Software entities should be open for extension, but closed for modification.
This means you should be able to add new functionality without changing existing code.
* In Go, interfaces can be used to achieve this.
You can define an interface and then extend its functionality by implementing the
interface with different types without modifying the original interface definition.
*/

type Saver interface {
	Save(data string)
}

type FileSaver struct{}

func (fs FileSaver) Save(data string) {
	// Implementation to save data to a file
}

type CloudSaver struct{}

func (cs CloudSaver) Save(data string) {
	// Implementation to save data to the cloud
}
