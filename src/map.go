package main

import "fmt"

func main() {
	/*
		Maps adalah Object pada bahasa pemograman lain yaitu terdiri dari {key : value}
	*/

	//##Looping pada Map

	// tanggal := map[string]int{
	// 	"januari":  12,
	// 	"february": 10,
	// 	"maret":    5,
	// 	"april":    26,
	// }

	// for key, val := range tanggal {
	// 	fmt.Println(key, "  \t:", val)
	// }

	//##Maps dan Slice

	mahasiswa := []map[string]int{
		{"nim": 17200435},
		{"nim": 1720345},
		{"nim": 17092394},
		{"nim": 17200523},
	}

	for _, nims := range mahasiswa {
		fmt.Print(nims["nim"], " ")
	}
}
