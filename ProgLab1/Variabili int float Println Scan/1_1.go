// Scrivere una applicazione Go che, dato in input un valore intero, lo elevi alla quarta potenza.

package main

import "fmt"

func main() {
	var a int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&a)
	fmt.Println(a, "elevato alla quarta diventa", a*a*a*a)
}
