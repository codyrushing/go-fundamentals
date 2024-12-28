package data

import "fmt"

type Instructor struct {
	// note that even properties need to be capital case to be public
	// lower case properties will be private
	Id        int
	FirstName string
	LastName  string
	Score     int
}

// this is how you add a method to the Instructor struct
func (i Instructor) Print() {
	fmt.Printf("%v %v (%d)", i.FirstName, i.LastName, i.Score)
}

func NewInstructor(firstName string, lastName string) Instructor {
	return Instructor{FirstName: firstName, LastName: lastName}
}
