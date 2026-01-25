package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	c := 20
	d := a + b*c

	fmt.Println(d)

	// Augmented assigment
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)
}
