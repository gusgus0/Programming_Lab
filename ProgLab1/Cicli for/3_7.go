/*
Scrivere un programma che legga un valore intero positivo e verifichi che sia composto di soli {0,1}, convertendolo nel corrispondente valore
"binario" (es. 100 -> 4). Nel caso in cui il numero contenga altre cifre, stampi un messaggio di errore.
*/

package main

import "fmt"

func main() {
	var bin, i, ris, conta, potenza int

	fmt.Print("Inserisci un numero in forma binaria: ")
	fmt.Scan(&bin)
	
	// per controllare che il numero inserito sia effettivamente un numero binario continuo a dividerlo per 10 e controllo che il 
	// resto sia sempre di solo 1 o 0 altrimenti stampo un messaggio di errore

	conta = 0
	ris = 0
	potenza = 1

	for bin >= 1 {
		if bin%10 != 1 && bin%10 != 0{
			fmt.Println("Errore: il numero inserito non è binario.")
			return
		} else {
			if bin%10 == 1 {
				for i = 0; i < conta; i++ {
					potenza = potenza * 2
				}
				ris = ris + potenza
				potenza = 1
				conta++
			} else if bin%10 == 0 {
				conta++
			}

			bin = bin / 10
		}
	}
	fmt.Println("In forma decimale diventa:", ris)
}
