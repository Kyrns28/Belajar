package main

import "fmt"

func sayHellow(name string) string {
	hello := "Hellowww " + name
	return hello
}

func main() {
	result := sayHellow("Rizkinur")
	fmt.Println(result)
}
