package main

import (
	"fmt"
)

func main() {

	var num1 int
	fmt.Println("Enter 1st Number: ")
	fmt.Scanln(&num1)

	var s string
	fmt.Println("Would You Like To [A]dd, [S]ubtract, [M]ultiply, [D]ivide? ")
	fmt.Scanln(&s)

	var num2 int
	fmt.Println("Enter 2nd Number: ")
	fmt.Scanln(&num2)

	switch s {
	case "A":
		sum := add(num1, num2)
		fmt.Println("Sum is", sum)
	case "S":
		diff := subtract(num1, num2)
		fmt.Println("Difference is", diff)
	case "M":
		product := multiply(num1, num2)
		fmt.Println("Product is", product)
	case "D":
		quotient := divide(num1, num2)
		fmt.Println("Quotient is", quotient)
	default:
		fmt.Println("Not valid input")
	}

}

func add(num1 int, num2 int) int {
	return int(num1) + int(num2)
}

func subtract(num1 int, num2 int) int {
	return num1 - num2
}

func multiply(num1 int, num2 int) int {
	return num1 * num2
}

func divide(num1 int, num2 int) float64 {
	if num2 == 0 {
		panic("Cannot divide by 0")
	}

	var quotient float64
	quotient = float64(num1) / float64(num2)
	return quotient
}

