package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("What is your current age?")
	var currentAge int
	fmt.Scanln(&currentAge)

	fmt.Println("At what age would you like to retire?")
	var retirementAge int
	fmt.Scanln(&retirementAge)

	yearsToRetirement := retirementAge - currentAge
	fmt.Printf("You have %d years left until you can retire. \n", yearsToRetirement)

	t := time.Now()
	fmt.Printf("It's %d, so you can retire in %d. \n", t.Year(), t.Year()+yearsToRetirement)
}
