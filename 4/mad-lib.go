package main

import "fmt"

func main() {
	fmt.Println("Enter a noun:")
	var noun string
	fmt.Scanln(&noun)

	fmt.Println("Enter a verb:")
	var verb string
	fmt.Scanln(&verb)

	fmt.Println("Enter an adjective:")
	var adjective string
	fmt.Scanln(&adjective)

	fmt.Println("Enter an adverb:")
	var adverb string
	fmt.Scanln(&adverb)

	fmt.Printf("Do you %s your %s %s %s? That's hilarious! \n", verb, adjective, noun, adverb)
}
