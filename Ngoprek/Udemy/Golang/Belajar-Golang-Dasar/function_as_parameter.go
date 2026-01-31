package main

import "fmt"

type Filter func(string) string // Wajib diisi jika func dijadikan type declaration

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

// Function with type declaration
func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hallooo", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	// contoh
	spamFilter(("Rizki"))
	sayHelloWithFilter("Rizki", spamFilter)

	filterr := spamFilter
	sayHelloWithFilter("Anjing", filterr)
}
