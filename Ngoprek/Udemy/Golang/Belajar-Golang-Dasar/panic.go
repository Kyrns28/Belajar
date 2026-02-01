package main

import "fmt"

func endApp() {
	fmt.Println("End APP")
}

func runApp(error bool) {
	// func tetap dijalankan diakhir saat func runApp selesai dijalankan walaupun error
	defer endApp()

	if error {
		// jika kode terjadi error akan ditampilkan beserta log nya
		panic("Ups error")
	}
}

func main() {
	runApp(true)
}
