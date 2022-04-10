package main

import "fmt"

func main() {

	//pointer adalah alamat memori pada suati vatiabel
	var nomorA int = 4
	var nomorB *int = &nomorA

	fmt.Println(*nomorB)
}
