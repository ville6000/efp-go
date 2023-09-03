package main

import "fmt"

func main() {
	fmt.Println("What is the first number?")
	var firstNumber int
	fmt.Scanln(&firstNumber)

	fmt.Println("What is the second number?")
	var secondNumber int
	fmt.Scanln(&secondNumber)

	fmt.Printf("%d + %d = %d \n", firstNumber, secondNumber, add(firstNumber, secondNumber))
	fmt.Printf("%d - %d = %d \n", firstNumber, secondNumber, subtract(firstNumber, secondNumber))
	fmt.Printf("%d * %d = %d \n", firstNumber, secondNumber, multiply(firstNumber, secondNumber))
	fmt.Printf("%d / %d = %d \n", firstNumber, secondNumber, divide(firstNumber, secondNumber))
}

func add(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func subtract(firstNumber int, secondNumber int) int {
	return firstNumber - secondNumber
}

func multiply(firstNumber int, secondNumber int) int {
	return firstNumber * secondNumber
}

func divide(firstNumber int, secondNumber int) int {
	return firstNumber / secondNumber
}
