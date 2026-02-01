package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	// jika menggunakan defer, func akan tetap dieksekusi walaupun terjadi error di function
	defer logging()

	fmt.Println("Run Application")
	// logging()
}

func main() {
	runApplication()
}
