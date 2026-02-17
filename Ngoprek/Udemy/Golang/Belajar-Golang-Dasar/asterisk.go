package main
import "fmt"

type Address struct {
	city, province, country string
}

func main(){
	var address1 = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 = &address1 // pointer
	address2.city = "Bandung"
	fmt.Println(address1) // ikut berubah menjadi bandung
	fmt.Println(address2) // berubah menjadi bandung

	// pointer ke struct Address
	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} // tidak akan mengubah nilai address1
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} 
	fmt.Println(address1)
	fmt.Println(address2)

}