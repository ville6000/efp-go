package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	var greeting string
	greeting = "Hello, " + name + ", nice to meet you!"

	if len(name) > 6 {
		greeting = "Haudihou, " + name + ", very nice to meet you!"
	}

	fmt.Println(greeting)
}
