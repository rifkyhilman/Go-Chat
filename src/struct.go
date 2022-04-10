package main

import "fmt"

/*
Struct adalah Object (Class) di  bahasa pemograman Golang karna Golang tidak memiliki Class layaknya bahasa pemograman OOP lainnya,
namun Golang Memiliki Type data layaknya object yaitu STRUCT
*/

type Mahasiswa struct {
	nama string
	usia int
}

func main() {
	// //Deklarasi Variabel object 1
	// var s1 Mahasiswa
	// s1.nama = "Rifki"
	// s1.usia = 20

	// //Deklarasi Variabel object 2
	// var s2 = Mahasiswa{}
	// s2.nama = "Dicky"
	// s2.usia = 22

	// //Deklarasi Variabel object 3
	// var s3 = Mahasiswa{
	// 	"Udin",
	// 	22,
	// }

	// //Deklarasi Variabel object 4
	// var s4 = Mahasiswa{
	// 	nama: "Muhamad",
	// 	usia: 22,
	// }

	// fmt.Println("Nama :", s1.nama)
	// fmt.Println("Usia :", s1.usia)

	// fmt.Println("Nama :", s2.nama)
	// fmt.Println("Nama :", s2.nama)

	// fmt.Println("Nama :", s3.nama)
	// fmt.Println("Nama :", s4.nama)

	//## Struct Anonymous
	var s1 = struct {
		Mahasiswa
		Gelar string
	}{}

	s1.nama = "Rifki"
	s1.usia = 20
	s1.Gelar = "Teknologi Informasi"

	fmt.Println(s1.nama)
	fmt.Println(s1.usia)
	fmt.Println(s1.Gelar)

}
