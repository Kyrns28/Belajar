package main

import "fmt"

func getFullName() (string, string) {
	return "Rizkinur", "Sepriana"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// menghiraukan return value
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
