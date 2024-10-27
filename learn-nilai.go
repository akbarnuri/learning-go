package main

import "fmt"

func main() {
	nilaiSaya := 90
	var grade string

	if nilaiSaya > 80 {
		grade = "A"
	} else if nilaiSaya > 70 {
		grade = "B"
	} else if nilaiSaya > 60 {
		grade = "C"
	} else if nilaiSaya > 50 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println("Nilai saya adalah : " + grade)
}
