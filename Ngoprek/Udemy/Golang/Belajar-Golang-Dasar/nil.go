// Nil biasa digunakan di beberapa tipe data seperti interface, function, map, slice , pointer dan channel
package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("Kinur")

	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {

		fmt.Println(data["name"])
	}
}
