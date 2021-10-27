/*
Scrivere un programma che stampa i primi n elementi della successione di fibonacci, con n scelto dall'utente.

fib(1) = fib(2) = 1
fib(n) = fib(n-1) + fib(n-2), se n > 2
*/

package main

import "fmt"

func main() {
	var n, primo, secondo, terzo, i int

	fmt.Print("Quanti elementi della successione di Fibonacci vuoi stampare? ")
	fmt.Scan(&n)

	primo = 1
	secondo = 1

	if n == 1 {
		fmt.Println(primo)
	}
	if n >= 2 {
		fmt.Println(primo)
		fmt.Println(secondo)
	}
	for i = 2; i < n; i++ {
		terzo = primo + secondo
		primo = secondo
		secondo = terzo
		fmt.Println(terzo)
	}
	/*
		1
		1
		1+1 = 2
		2+1 = 3
		3+2 = 5
		5+3 = 8
		8+5 = 13
		13+8 = 21
		21+13 = 34


	*/

}
