package main

import (
	"fmt"
	"strings"
)

//##deklarasi Struct
type Mahasiswi struct {
	nama   string
	nourut int
}

func (s Mahasiswi) SayHello() {
	fmt.Println("Halo", s.nama)
}

func (s Mahasiswi) getNameAt(i int) string {
	return strings.Split(s.nama, " ")[i-1]
}

func main() {
	var s1 = Mahasiswi{nama: "udin", nourut: 21}
	s1.SayHello()

	var name = s1.getNameAt(1)
	fmt.Println("Nama panggilan :", name)
}
