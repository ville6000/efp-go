package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	employees := []string{"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen", "Jeremy Goodwin"}
	printState(employees)

	for {
		name := getRemoveableEmployee()

		if slices.Contains(employees, name) {
			employees = deleteEmployee(employees, name)
			break
		} else {
			fmt.Println("Invalid name")
		}
	}

	printState(employees)
}

func getRemoveableEmployee() string {
	fmt.Println("Enter employee name to remove:")
	var name string
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimRight(name, "\n")
}

func printState(employees []string) {
	fmt.Printf("There are %d employees: \n", len(employees))
	fmt.Println(employees)
}

func deleteEmployee(employees []string, employeeToRemove string) []string {
	return slices.DeleteFunc(employees, func(name string) bool {
		return name == employeeToRemove
	})
}
