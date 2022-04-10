package main

import "fmt"

func main() {
	//## Fungsi tanpa Return

	// names := []string{"John", "Wick"}
	// printMessage("halo", names)

	//## Fungsi dengan Return

	// kalian := perkalian(5, 5)

	// fmt.Println(kalian)

	//## Fungsi Variadic ( parameter tak terbatas)
	// var rata = kalkulasi(1, 3, 2, 3, 4, 5, 6)
	// var hasil = fmt.Sprintf("Rata-rata : %.2f", rata)
	// fmt.Println(hasil)

	//##Fungsi Clouser yaitu function di dalam variabel

	var ConClouser = func(x, y, z int) int {
		var perkalian = x * z
		// pembagian := y / x

		return perkalian
	}

	var Perkalian = ConClouser(5, 20, 2)

	fmt.Println("Hasil dari perkalian :", Perkalian)
}

//## Fungsi tanpa Return

// func printMessage(message string, arr []string) {
// 	namaString := strings.Join(arr, " ")

// 	fmt.Println(message, namaString)
// }

//## Fungsi dengan Return

// func perkalian(a, b int) int {

// 	var kali = a * b

// 	return kali
// }

//## Fungsi Variadic ( parameter tak terbatas)

// func kalkulasi(nomer ...int) float64 {

// 	var total int = 0

// 	for _, no := range nomer {
// 		total += no
// 	}

// 	var tengah = float64(total) / float64(len(nomer))

// 	return tengah

// }
