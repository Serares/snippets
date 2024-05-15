package SOLID

// Single Responsability Principle
/*
* A class or module should have one, and only one,reason to change.
This means that the class should have only one job or responsibility.
* In Go, rather than classes, you typically work with structs and interfaces.
A struct should be designed to serve a single functionality.
For example, a struct handling user information
should only manage user data and not take on additional
responsibilities like managing user sessions.
*/
type User struct {
	ID    string
	Name  string
	Email string
}

func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}
