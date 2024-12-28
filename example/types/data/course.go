package data

import "fmt"

type Duration float32

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor // Instructor
}

func (c Course) SignUp() bool {
	return true
}

// this is like implementing a `toString()` method on a class
// controls how Course is rendered to a string (eg. fmt.Print(someCourse))
func (c Course) String() string {
	return fmt.Sprintf(
		"%v --- %v",
		c.Name, c.Instructor.FirstName)
}
