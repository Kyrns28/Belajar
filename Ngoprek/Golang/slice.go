package main

import (
	"fmt"
)

// example slice
// func sliceGo() {
// 	mySlice1 := []int{}
// 	fmt.Println(len(mySlice1))
// 	fmt.Println(cap(mySlice1))
// 	fmt.Println(mySlice1)

// 	mySlice2 := []string{"Go", "Slices", "are", "Powerfull"}
// 	fmt.Println(len(mySlice2))
// 	fmt.Println(cap(mySlice2))
// 	fmt.Println(mySlice2)
// }

// Example create a slice from array
// func sliceGo() {
// 	arr1 := [6]int{10, 11, 12, 13, 14, 15}
// 	mySlice := arr1[2:4]

// 	fmt.Printf("MySlice = %v\n", mySlice)
// 	fmt.Printf("length = %d\n", len(mySlice))
// 	fmt.Printf("Capacity = %d\n", cap(mySlice))
// }

// With function make
// func sliceGo() {
// 	mySlice1 := make([]int, 5, 10)
// 	fmt.Printf("mySlice1 = %v\n", mySlice1)
// 	fmt.Printf("length = %d\n", len(mySlice1))
// 	fmt.Printf("capacity = %d\n", cap(mySlice1))

// 	mySlice2 := make([]int, 5)
// 	fmt.Printf("mySlice2 = %v\n", mySlice2)
// 	fmt.Printf("length = %d\n", len(mySlice2))
// 	fmt.Printf("capacity = %d\n", cap(mySlice2))
// }

// Slice menggunakan array (tanpa make)
// func sliceTanpaMake() {
// 	arr1 := [6]int{10, 11, 12, 13, 14, 15} // array asli
// 	mySlice := arr1[2:5]                   // meminjam sebagian array

// 	fmt.Println("Slice :", mySlice)      // output : Slice : [12 13]
// 	fmt.Println("Len   :", len(mySlice)) // output : Len   : 2
// 	fmt.Println("Cap   :", cap(mySlice)) // output : Cap   : 4

// 	mySlice[2] = 99
// 	fmt.Println(mySlice) // output : [12 13 99]
// }

// slice menggunakan make
// func slicePakaiMake() {
// 	mySlice := make([]int, 2, 4)

// 	fmt.Println("Slice :", mySlice)      // output : Slice : [0 0]
// 	fmt.Println("Len   :", len(mySlice)) // output : Len   : 2
// 	fmt.Println("Cap   :", cap(mySlice)) // output : Cap   : 4
// 	fmt.Println()

// 	mySlice[1] = 99
// 	fmt.Println(mySlice) // output : [0 99]
// 	fmt.Println()

// 	mySlice = append(mySlice, 100, 200, 300)
// 	fmt.Println(mySlice)                 // output [0 99 100 200 300]
// 	fmt.Println("Len   :", len(mySlice)) // output Len   : 5
// 	fmt.Println("Cap   :", cap(mySlice)) // output Cap   : 8

// 	fmt.Println()
// 	mySlice2 := make([]int, 0, len(mySlice))
// 	mySlice2 = append(mySlice2, mySlice...)
// 	fmt.Println(mySlice2) // [0 99 100 200 300]

// 	Slice1 := make([]string, 2, 5)
// 	Slice1[0] = "Apel"
// 	Slice1[1] = "Jeruk"
// 	fmt.Println(Slice1) // [Apel Jeruk]

// 	Slice2 := make([]string, len(Slice1))
// 	Slice2 = append(Slice2, Slice1...)
// 	fmt.Println("Slice 2 adalah : ", Slice2) // Slice 2 adalah :  [Apel Jeruk]

// 	Slice2 = append(Slice2, "melon")
// 	fmt.Println("Slice 2 yang terbaru adalah : ", Slice2) // Slice 2 yang terbaru adalah :  [Apel Jeruk melon]
// }

// Menggunakan Copy

func sliceCopy() {
	Buah := make([]string, 0, 5)
	Buah = append(Buah, "Jeruk", "Apel")
	fmt.Println(Buah)

	salinBuah := make([]string, len(Buah))
	copy(salinBuah, Buah)
	fmt.Println(salinBuah)

	salinBuah = append(salinBuah, "Melon")
	fmt.Println(salinBuah)
}
