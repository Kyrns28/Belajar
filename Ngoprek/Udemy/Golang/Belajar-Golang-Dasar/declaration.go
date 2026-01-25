package main

import "fmt"

func main() {
	type noKTP string

	var ktpKinur noKTP = "1111111111111"

	var contoh string = "22222222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(ktpKinur)
	fmt.Println(contohKTP)
}
