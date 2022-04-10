package main

import "fmt"

func main() {

	//##Array biasa

	// mahasiswa := [3]string{
	// 	"Ando",
	// 	"Haikal",
	// 	"Ahmad",
	// }

	// fmt.Println(mahasiswa)

	//##Array Tanpa Jumlah Element

	// banks := [...]int{
	// 	100,
	// 	90,
	// 	2000,
	// 	1023,
	// }

	// fmt.Println(banks)

	//##Array Multidimensi

	// nomerTogel := [2][3]int{
	// 	{3, 2, 1},
	// 	{9, 8, 7},
	// }

	// fmt.Println(nomerTogel)

	//##Array + Looping

	sayuran := [4]string{
		"tomat",
		"Toge",
		"Terong",
		"Bawang",
	}

	// for i := 0; i < len(sayuran); i++ {
	// 	fmt.Println("element:", i, sayuran[i])
	// }

	//Menampilkan indeks saja
	// for i, _ := range sayuran {
	// 	fmt.Println(i)
	// }

	//Menampilkan Element saja
	for _, sayur := range sayuran {
		fmt.Println("Nama Buah", sayur)
	}
}
