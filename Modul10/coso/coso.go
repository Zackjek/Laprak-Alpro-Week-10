package main

import "fmt"

func main() {
	var a, b, hasil float64
	fmt.Scan(&a, &b)
	if b != 0 {
		hasil = a / b
		fmt.Println("Hasil Pembagian Adalah", hasil)

	}else {
		fmt.Println("Variabel b Bernilai 0")
	}
	fmt.Print("Program selesai")
}