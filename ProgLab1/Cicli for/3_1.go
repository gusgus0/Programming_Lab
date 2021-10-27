/*
Scrivere un programma che stampa i primi n numeri triangolari, con n scelto dall'utente.
*/

package main

import "fmt"

func main() {
	var n, t, i int
	fmt.Print("n? ")
	fmt.Scan(&n)
	// stampiamo i primi n numeri naturali
	// l'i-esimo numero triangolare è 1 + 2 + ... i
	t = 0
	for i = 1; i <= n; i++ {
		t = t + i
		fmt.Println(t)
	}
}
