package main

import "fmt"

func main() {
    var b int
    fmt.Print("bilangan : ")
    fmt.Scan(&b)
    
    fmt.Print("Faktor : ")
    jumlahFaktor := 0
	 for i := 1; i <= b; i++ {
        if b % i == 0 {
            fmt.Print(" " ,i)
            jumlahFaktor++
        }
    }
	
	
	fmt.Println()


    if jumlahFaktor == 2 {
        fmt.Println( "prima : true")
    } else {
        fmt.Println( "prima : false")
    }
}
