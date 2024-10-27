package main

import "fmt"

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func calculateTaxWithName(price float32) (cityTax float32, stateTax float32) {
	cityTax = price * 0.09
	stateTax = price * 0.02
	return
}

// *int says that it will be accepting a pointer to an int, not an int as a value. it is a different type
func birthday(age *int) {
	if *age > 140 {
		panic("Too old!!!")
	}
	defer fmt.Printf("age is %v\n", *age)
	// *age says to get the value of the age reference, this is important. trying to do `age += 1` will fail because `age` is a reference, not an int
	// you need to use *age to get the value, which is here as an int
	*age += 1
}

func main() {
	cityTax, stateTax := calculateTax(100)
	fmt.Println(cityTax)
	fmt.Println(stateTax)
	age := 22
	// &age is the reference to the age variable. gets the reference for a variable
	birthday(&age)
	fmt.Println(age)

	endOfGame := false
	for !endOfGame {
		fmt.Println("lets end the game now")
		endOfGame = true
	}

}
