package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Bandung" // pass by value

	// fmt.Println(address1) // Tidak berubah
	// fmt.Println(address2) // City berubah menjadi bandung

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // Pointer

	address2.City = "Bandung"

	fmt.Println(address1) // City berubah menjadi Bandung karena address2 pointer ke address1
	fmt.Println(address2) // City berubah menjadi bandung
}
