package data

import "time"

type Workshop struct {
	// multiple type "inheritance"
	// if there are name collisions
	Course     // this type will be embedded into Workshop, all of its properties now exist on Workshop
	Instructor // same with this
	Date       time.Time
}

func (w Workshop) SignUp() bool {
	return true
}

func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	// Course embedded type also has a name, how to access it?
	// w.Course.Name <= you can just disambiguate it this way
	w.Name = name
	w.Instructor = instructor
	return w
}
