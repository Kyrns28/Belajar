package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Ini perulangan yang ke ", counter)
		counter++
	}

	// for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Ini perulangan yang ke ", counter)
	}

	fmt.Println("Selesai")
	fmt.Println()

	names := []string{"Rizki", "Nur", "Sepriana"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println()
	// for range dengan index
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
	fmt.Println()

	// for range tanpa index
	for _, name := range names {
		fmt.Println(name)
	}
}
