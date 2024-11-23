package main

import "fmt"

func main() {
	var usia int
	var kk bool
	fmt.Scan(&usia, &kk)
	if usia >= 17 && kk {
		fmt.Print("Bisa membuat KTP")

	}else {
		fmt.Print("Belum bisa membuat KTP")
	}
}