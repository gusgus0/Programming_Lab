/*
Scrivere un programma che verifichi se una data nel formato
gg mm aa sia valida, visualizzando nel caso un messaggio d'errore.
*/

package main

import "fmt"

func main() {
	var gg, mm, aa int

	fmt.Print("Inserire data nel formato gg mm aa: ")
	fmt.Scan(&gg, &mm, &aa)

	if gg > 31 || mm > 12 {
		fmt.Println("ERRORE: si è inserito un valore non valido")
	} else if (gg > 28 && mm == 2) || (gg > 30 && mm == 4) || (gg > 30 && mm == 6) || (gg > 30 && mm == 9) || (gg > 30 && mm == 11) {
		fmt.Println("ERRORE: si è inserito un valore non valido per il mese scelto")
	}
}
