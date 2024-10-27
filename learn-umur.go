package main

import "fmt"

func main() {
	umurSaya := 20
	umursayatahunDepan := 1
	umursayalimaTahun := 5
	umurKakak := 30
	umurkakaktahunDepan := 1
	umurkakaklimaTahun := 5
	fmt.Printf("Umur saya sekarang adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak sekarang adalah : %d\n", umurKakak)

	// tahun depan
	umurSaya += umursayatahunDepan
	umurKakak += umurkakaktahunDepan

	fmt.Printf("Umur saya tahun depan adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak tahun depan adalah : %d\n", umurKakak)

	// 5 tahun lagi
	umurSaya += umursayalimaTahun
	umurKakak += umurkakaklimaTahun

	fmt.Printf("Umur saya 5 tahun lagi adalah : %d\n", umurSaya)
	fmt.Printf("Umur kakak 5 tahun lagi adalah : %d\n", umurKakak)

}
