package main

import "fmt"

func main(){
	var a = 10
	var b = 10
	c := 2
	d := a+b*c

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)

	// assignment operator
	var e = 10
	e += 5
	fmt.Println(e)

	// unary operator
	f := 10
	f++ // f = f+1
	fmt.Println(f)
	f-- // f=f-1
	fmt.Println(f)
}