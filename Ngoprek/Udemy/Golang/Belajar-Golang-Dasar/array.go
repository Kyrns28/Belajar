package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Rizki"
	names[1] = "Nur"
	names[2] = "Sepriana"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])


	// deklarasi array secara langsung
	var value = [...]int {
		80,
		90,
		20,
		40,
	}

	fmt.Println(value) // output 80,90,0 (karena nilai default array int kosong itu 0)
	fmt.Println("Index ke -0 = ", value[0]) // output 80
	fmt.Println("Index ke -1 = ", value[1]) // output 90
	fmt.Println("Index ke -2 = ",value[2]) // output 0

	fmt.Println("Total nilai array ada = ", len(value))
}
