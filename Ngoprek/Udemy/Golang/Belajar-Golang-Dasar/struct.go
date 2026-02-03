package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

// Method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.name)
}

func main() {
	var rizki Customer
	fmt.Println(rizki)

	rizki.name = "Rizki Nur"
	rizki.address = "Depok"
	rizki.age = 24

	fmt.Println(rizki)
	fmt.Println(rizki.name)
	fmt.Println(rizki.address)
	fmt.Println(rizki.age)

	joko := Customer{
		name:    "Joko aja",
		address: "Jakarta",
		age:     28,
	}
	fmt.Println(joko)

	Hermawan := Customer{"Hermawan Aja", "Bogor", 25}
	fmt.Println(Hermawan)

	// Memanggil method
	joko.sayHello("Kinur")
}
