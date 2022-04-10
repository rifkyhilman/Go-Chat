package main

import "fmt"

//##deklarasi Struct
type Mahasiswi struct {
	nama   string
	nourut int
}

func (s Mahasiswi) SayHello() {
	fmt.Println("halo", s.nama)
}

func main() {

}
