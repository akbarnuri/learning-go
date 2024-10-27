package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukan angka: ")
	fmt.Scan(&number)

	//menggunakan modulus dan percabangan untuk mengecek ganjul atau genap

	if number%2 == 0 {
		fmt.Printf("%d adalah bilangan genap\n", number)
	} else {
		fmt.Printf("%d adalah bilangan ganjil\n", number)
	}
}
