// Scrivere un programma che legga un valore intero e stampi "fizz" se il numero è multiplo di 3,
// "buzz" se il numero è multiplo di 5, "fizz buzz" se multiplo sia di 3 sia di 5, una riga vuota altrimenti

package main

import "fmt"

func main() {
	var val int

	fmt.Print("Inserisci un valore: ")
	fmt.Scan(&val)

	if val%3 == 0 {
		fmt.Print("fizz")
	}
	if val%5 == 0 {
		fmt.Print("buzz")
	} else {
		fmt.Println()
	}
}
