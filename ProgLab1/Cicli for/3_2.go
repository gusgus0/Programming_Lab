/*
Scrivere un programma che stampa la somma dei primi n numeri dispari, con n scelto dall'utente
*/

package main

import "fmt"

func main() {
	var i, sum, n int
	fmt.Print("Quanti numeri dispari vuoi sommare? ")
	fmt.Scan(&n)

	sum = 0
	for i = 0; i < n; i++ { // NB: i++ e la stessa cosa che scrivere i += 1 (i = i + 1)
		sum = sum + (2*i + 1)
	}

	fmt.Println("La somma dei primi", n, "numeri dispari è pari a ", sum)
}
