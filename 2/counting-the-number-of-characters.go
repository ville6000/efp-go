package main

import "fmt"

func main() {
    fmt.Println("What is the input string?")
    var input string
    fmt.Scanln(&input)

    fmt.Printf("%s has %d characters. \n", input, len(input))
}
