package main

import "fmt"

func main() {
	umur := 23
	var hasKTP bool

	if umur >= 17 {
		hasKTP = true
	} else {
		hasKTP = false
	}

	fmt.Printf("Apakah anda sudah mempunyai ktp ? %v", hasKTP)
}
