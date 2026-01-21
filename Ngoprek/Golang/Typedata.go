package main

import (
	"fmt"
)

func showTypeData() {
	var a bool = true    // Boolean
	var b int = 5        // integer
	var c float64 = 3.14 //floating point number
	var d string = "Hi"  // String

	e := "Hellow end" // untyped without var

	fmt.Println("Boolean : ", a)
	fmt.Println("Integer : ", b)
	fmt.Println("Float : ", c)
	fmt.Println("String : ", d)
	fmt.Println()
	fmt.Println(e)
}
