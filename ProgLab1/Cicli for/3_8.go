/*
Scrivere un programma che legga una sequenza di coppie di valori interi, interpretate come frazioni, e ne calcoli
la somma esatta, cioè come frazione. Una "frazione" con denominatore 0 denoti la fine dell'input
*/

package main

import "fmt"

func main() {
	var num, den int
	var sommaNum, sommaDen int = 0, 0

	fmt.Println("Inserisci una sequenza di frazioni: ")

	for {
		fmt.Scan(&num, &den)

		if den == 0 {
			break
		}

		sommaNum = sommaNum + num
		sommaDen = sommaDen + den
	}

	

	fmt.Println("La somma è: ", sommaNum, "/", sommaDen)
}