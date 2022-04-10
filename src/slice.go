package main

import "fmt"

func main() {
	//##slice kembaran nya array
	sayur := []string{"tomat", "cabe", "kangkung"}
	// //fungsi append untuk menambahkan data ke dalam slice
	sayuran := append(sayur, "bayam", "bawang")

	fmt.Println(sayur[1])
	fmt.Println(sayuran[4])

}
