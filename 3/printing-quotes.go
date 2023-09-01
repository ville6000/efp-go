package main

import "fmt"

func main() {
	var author string
	var quote string

	fmt.Println("What is the quote?")
	fmt.Scanln(&quote)
	fmt.Println("Who said it?")
	fmt.Scanln(&author)

	fmt.Printf("%s says, \"%s\" \n", author, quote)
}
