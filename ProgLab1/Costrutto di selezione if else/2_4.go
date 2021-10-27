// Scrivere un programma che esprima un voto (un valore intero) in trentesimi, secondo la seguente scala: se > 30 "out of range",
// tra 28 e 30 "eccellente", tra 25 e 27 "buono", tra 21 e 24 "discreto", tra 18 e 20 "sufficiente", altrimenti "insufficiente".
// Il programma legge un voto numerico e stampa il messaggio corrispondente

package main

import "fmt"

func main() {
	var voto uint

	fmt.Print("Inserire il voto: ")
	fmt.Scan(&voto)

	if voto > 30 {
		fmt.Println("out of range")
	} else if voto >= 28 {
		fmt.Println("Eccellente!")
	} else if voto >= 25 {
		fmt.Println("Buono")
	} else if voto >= 21 {
		fmt.Println("Discreto")
	} else if voto >= 18 {
		fmt.Println("Sufficiente")
	} else {
		fmt.Println("Insufficiente")
	}
}
