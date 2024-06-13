package main

import "fmt"

type myStr string

type myStruck struct {
	kode int
	nama string
}

func main() {
	var a myStr = "Custome str"
	println(a)

	var b produk = produk{kode: 2, nama: "buku"}
	//b.kode = 1
	//b.nama = "tas"
	fmt.Println(b)
	b.cetak()
	a.cetak("Hallo")
}

func (n myStr) cetak(str string) {
	println(str, n)
}

func (p produk) cetak() {
	println(p.kode, p.nama)
}
