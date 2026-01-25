package main
import "fmt"

func main(){
	names := [...]string {"Rizki", "Nur", "Sepriana", "Tes1", "Tes2", "Tes3"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	fmt.Println()
	// 
	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:] // output : sabtu, minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days) // index 5 dan index 6 akan kerubah mengikuti daySlice
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Libur Baru") // menambah nilai di index terakhir
	fmt.Println(daySlice2) // output : sabtu baru, minggu baru, libur baru

	daySlice2[0] = "sabtu lama"
	fmt.Println(daySlice2)
	fmt.Println()

	fmt.Println(days)
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
// | Baris          | Kenapa                                                  |
// | -------------- | ------------------------------------------------------- |
// | **Line 26–27** | Slice masih pakai array `days`, jadi array ikut berubah |
// | **Line 28**    | `append` bikin **array baru** karena cap tidak cukup    |
// | **Line 32**    | Ubah slice baru → **days tidak terpengaruh**            |
	fmt.Println()

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Rizki"
	newSlice[1] = "Nur"
	// newSlice[2] = "Sepriana" // akan error karena len slice hanya 2, jika menambahkan nilai harus menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println()

	newSlice2 := append(newSlice, "Sepriana")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
	fmt.Println()

	newSlice2[0] = "Kosong"
	fmt.Println(newSlice) // output  : Kosong , Nur (karena masih dalam array yang sama sehingga ikut berubah)
	fmt.Println(newSlice2) // output : kosong, Nur, Sepriana (karena masih dalam array yang sama sehingga ikut berubah)
	fmt.Println()

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
	toSlice[0] = "tess"
	fmt.Println(toSlice)
	fmt.Println(fromSlice)

	// Catatan tambahan perbedaan array dan slice
	// iniArray := [...]int{1,2,3,4,5}
	// iniSlice := []int{1,2,3,4,5}
}