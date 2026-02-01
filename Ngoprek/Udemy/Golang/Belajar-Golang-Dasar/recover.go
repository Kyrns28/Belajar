package main

import "fmt"

func endApp() {
	fmt.Println("end APP")

	message := recover()
	fmt.Println("terjadi panic dengan pesan ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ups error")
	}
}

func main() {
	runApp(true)
}
