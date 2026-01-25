package main
import "fmt"

func main(){
	name := Rizkinur

	if name == "Eko"{
		fmt.Println("Halo Eko")
	} else if name == "Rizkinur"{
		fmt.Println("Halo Rizkinur")
	} else if name == "Joko"{
		fmt.Println("Halo Joko")
	} else {
		fmt.Println("Boleh kenalan ?")
	}

	// IF short statement
	length := len(name)
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}