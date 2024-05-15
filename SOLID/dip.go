package SOLID

// Dependency Inversion Principle
/*
* High-level modules should not depend on low-level modules.
Both should depend on abstractions (e.g., interfaces, not concretions).
* Instead of directly depending on concrete implementations, your code should depend on interfaces.
*/
type Database interface {
	Connect()
}

type MySQLDatabase struct{}

func (mdb MySQLDatabase) Connect() {
	// MySQL connection logic
}

type Service struct {
	DB Database
}

func NewService(db Database) *Service {
	return &Service{DB: db}
}
