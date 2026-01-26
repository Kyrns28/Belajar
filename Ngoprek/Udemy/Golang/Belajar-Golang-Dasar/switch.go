package main
import "fmt"

func main(){
	name := "Rizkinur"

	switch name{
	case "Eko":
		fmt.Println("Halo Eko")
	case "Budi":
		fmt.Println("Halo Budi")
	case "Rizkinur":
		fmt.Println("Halo Rizkinur")
	default:
		fmt.Println("Kenalan dulu")
	} 

	// switch short statement
	switch length := len(name); length >5 {
		case true :
			fmt.Println("Terlalu panjang")
		case false:
			fmt.Println("Nama sudah benar")
	}

	name = "RizkinurRizkinur"
	// Switch tanpa kondisi
	length := len(name)
	switch{
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5 :
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	} 
}