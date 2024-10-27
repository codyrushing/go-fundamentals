package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Println("Calculator 1.0")
	fmt.Println("==============")

	fmt.Println("Which operation do you want to perform (add, subtract, multiply, divide)")
	// reads from stdin. here we actually want to change the value of operation, so we pass a reference
	fmt.Scanf("%s", &operation)
	fmt.Println("First number:")
	fmt.Scanf("%d", &num1)
	fmt.Println("Second number:")
	fmt.Scanf("%d", &num2)

	switch operation {
	case "add":
		fmt.Println(num1 + num2)
	default:
		fmt.Printf("Invalid operation %s", operation)
	}
}
