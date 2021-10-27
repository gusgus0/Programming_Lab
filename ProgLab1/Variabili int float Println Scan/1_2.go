// Scrivere una applicazione Go che, dati tre numeri, ne calcoli la somma

package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Inserisci tre numeri interi: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(a, "+", b, "+", c, "=", a+b+c)
}
