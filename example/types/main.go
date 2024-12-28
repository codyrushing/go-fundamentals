package main

import (
	"fmt"

	"example.com/types/scratch/data"
)

func main() {
	instructor := data.Instructor{Id: 5, LastName: "Firtman"}
	// you can add props after construction
	instructor.FirstName = "Max"
	instructor.Print()

	fmt.Println()
	kyle := data.NewInstructor("Kyle", "Simpson")
	kyle.Print()

	fmt.Println()
	goCourse := data.Course{Id: 1, Name: "Go Fundamentals", Instructor: kyle}
	fmt.Print(goCourse)
	fmt.Println()

	swiftWorkshop := data.NewWorkshop("Swift with iOS", instructor)
	fmt.Printf("%v", swiftWorkshop)

	// if the type here were [2]data.Course, it would not compile
	// because Workshop and Course are different types even though Course
	// is embedded into Workshop
	// instead we use an interface that both types successfully implement
	var courses [2]data.SignupAble
	courses[0] = goCourse
	courses[1] = swiftWorkshop

	for _, course := range courses {
		fmt.Println(course)
	}
}
