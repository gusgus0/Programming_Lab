/*
Scrivere un programma che stampa la potenza n-esima (n intero positivo) di un numero x.
*/

package main

import "fmt"

func main() {
	var x, n, i int
	var pow float64

	fmt.Print("base? ")
	fmt.Scan(&x)
	fmt.Print("esponente? ")
	fmt.Scan(&n)

	pow = 1

	for i = 0; i < n; i++ {
		pow = pow * float64(x)
	}

	fmt.Println(x, "elevato a", n, "=", pow)
}
