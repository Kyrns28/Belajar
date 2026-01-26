package main

import "fmt"

func getGoodbye(name string) string {
	return "goodbye " + name
}

func main() {
	contoh := getGoodbye

	fmt.Println(contoh("Rizkinur"))
}
