/*
Scrivere un programma che legge una serie di valori interi positivi e dopo ogni lettura stampa "crescente"
se il valore letto è maggiore o uguale al precedente, "decrescente" altrimenti. Il programma si ferma quando
il valore letto è negativo.
*/

package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Inserire n: ")
	fmt.Scan(&a)

	for {
		fmt.Print("Inserire n: ")
		fmt.Scan(&b)
		if b < 0 {
			break
		}
		if b >= a {
			fmt.Println("crescente")
			a = b
		} else if a > b {
			fmt.Println("decrescente")
		}
	}
}
