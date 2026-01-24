package main
import "fmt"

func main(){
	// Konversi tipe data I
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32) // konversi nilai32 menjadi int64
	var nilai16 int16 = int16(nilai32) // output akan berubah minus karena 32768 tidak termasuk cakupan int16 (max 32767)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Konversi tipe data II
	var name = "Rizki Nur Sepriana"
	var e uint8 = name[0]
	var eString = string(e)
	
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}