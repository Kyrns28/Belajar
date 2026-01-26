package main

import "fmt"

func getCompletedNames() (firstName, middleName, lastName string) {
	firstName = "Rizki"
	middleName = "Nur"
	lastName = "Sepriana"

	return firstName, middleName, lastName
}

func main() {

	a, b, c := getCompletedNames()
	fmt.Println(a, b, c)
}
