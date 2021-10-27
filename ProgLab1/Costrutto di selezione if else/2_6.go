/*
Scrivere un programma che, dati due appuntamenti giornalieri (un appuntamento è descritto da due coppie di interi
[hh,mm]-con hh 0..24 e mm 1 ..60- che ne rappresentano l'inizio e la fine) verifichi se siano (anche parzialmente)
sovrapposti, visualizzando nel caso un messaggio. Come unico controllo sui dati di input si accerti che l'orario
di inizio di un appuntamento preceda quello di fine, visualizzando un messaggio di errore in caso contrario.
*/

package main

import "fmt"

func main() {
	var hh1_i, mm1_i, hh1_f, mm1_f, hh2_i, mm2_i, hh2_f, mm2_f int = 0, 0, 0, 0, 0, 0, 0, 0

	// controllo che l'inizio del primo appuntamento non sia superiore a quello della sua fine
	for ((hh1_f*60)+mm1_f-(hh1_i*60)+mm1_i) <= 0 || ((hh2_f*60)+mm2_f-(hh2_i*60)+mm2_i) <= 0 {
		fmt.Print("Inserire ora (hh) e minuti (mm) dell'inizio del primo appuntamento: ")
		fmt.Scan(&hh1_i, &mm1_i)
		fmt.Print("Inserire ora (hh) e minuti (mm) della fine del primo appuntamento: ")
		fmt.Scan(&hh1_f, &mm1_f)

		fmt.Print("Inserire ora (hh) e minuti (mm) dell'inizio del secondo appuntamento: ")
		fmt.Scan(&hh2_i, &mm2_i)
		fmt.Print("Inserire ora (hh) e minuti (mm) della fine del secondo appuntamento: ")
		fmt.Scan(&hh2_f, &mm2_f)

		if ((hh1_f*60)+mm1_f-(hh1_i*60)+mm1_i) <= 0 || ((hh2_f*60)+mm2_f-(hh2_i*60)+mm2_i) <= 0 {
			fmt.Println("ERRORE: l'orario di inizio inserito viene dopo l'orario di fine!")
		}
	}

	// controlla se gli orari di inizio coincidono
	if hh1_i == hh2_i {
		fmt.Println("Gli impegni sono sovrapposti!")
	} else if hh1_f > hh2_i && hh1_f < hh2_f {
		fmt.Println("Gli impegni sono parzialmente sovrapposti!") // se l'orario di fine del primo appuntamento è compreso nella durata del secondo appuntamento
	} else if hh2_f > hh1_i && hh2_f < hh1_f {
		fmt.Println("Gli impegni sono parzialmente sovrapposti!")
	} else {
		fmt.Println("Gli impegni non si sovrappongono.")
	}

}
