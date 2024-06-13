package main

func main(){
	i := 10
	p := &i
	q := &i
	//println(*p)
	//*p =20
	//println(i)
	println("i", i) //memanggil nilai pada i  = 10
	println("*p", *p) //* atau pointer ini memanggil nilai pada i = 10
	println("q", *q) // memanggil alamat pada nilai i = 0xc000047f30

	i = 20
	println("i", i)
	println("*p", *p)
	println("q", *q)

	*p = 100
	println("i", i)
	println("*p", *p)
	println("q", *q)
}