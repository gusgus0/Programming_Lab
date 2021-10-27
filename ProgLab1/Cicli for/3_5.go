/*
Scrivere un programma che stampa il fattoriale di un numero naturale x.
Si ricorda che 0! = 1, e, se x > 0, x! = x * (x-1) * (x-2) * ... * 1
*/

package main

import "fmt"

func main() {
	var n, res, i int

	fmt.Print("Inserire il numero di cui si vuole sapere il suo fattoriale: ")
	fmt.Scan(&n)

	res = 1

	for i = 0; i < n; i++ {
		res = res * (n - i)
	}

	fmt.Println("Il fattoriale di", n, "risulta:", res)
}
