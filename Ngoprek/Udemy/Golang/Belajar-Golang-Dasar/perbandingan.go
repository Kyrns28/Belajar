package main

import "fmt"

func main(){
	var name1 = "Rizki"
	var name2 = "Rizki"
	name3 := "Nur"
	
	var result1 bool = name1 == name2
	var result2 bool = name1 != name3
	fmt.Println(result1)
	fmt.Println(result2)
}