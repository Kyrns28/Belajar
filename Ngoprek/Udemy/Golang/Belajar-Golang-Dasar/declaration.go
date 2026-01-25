package main
import "fmt"

func main(){
	type noKTP string

	var ktpRizkinur noKTP = "11111111111"
	
	var contoh string = "22222222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(ktpRizkinur)
	fmt.Println(contohKTP)
}