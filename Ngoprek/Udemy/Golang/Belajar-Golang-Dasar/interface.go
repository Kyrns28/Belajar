package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("Helloww", value.GetName())
}

// membuat struct
type Person struct {
	Name string
}

// Membuat Method
func (person Person) GetName() string {
	return person.Name
}

// Membuat struct
type Animal struct {
	Name string
}

// Membuat method
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	person := Person{Name: "Rizki"}
	sayHello(person)

	animal := Animal{Name: "Kuda"}
	sayHello(animal)
}
