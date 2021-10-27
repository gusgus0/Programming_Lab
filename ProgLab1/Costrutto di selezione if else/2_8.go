/*
Scrivere un programma che determini la posizione relativa di punto (x,y) dato rispetto ad una retta
data (m,q) data (sopra, sotto, appartiene)
*/

package main

import "fmt"

func main() {
	var m, q int = 2, 3 // y = 2x + 3
	var x, y int

	fmt.Print("Inserire la coordinata x del punto: ")
	fmt.Scan(&x)
	fmt.Print("Inserire la coordinata y del punto: ")
	fmt.Scan(&y)

	if (m*x)+q == y {
		fmt.Println("appartiene")
	} else if (m*x)+q > y {
		fmt.Println("sotto")
	} else if (m*x)+q < y {
		fmt.Println("sopra")
	}
}
