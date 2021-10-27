// Scrivere una applicazione Go che converta una durata (intero positivo) espressa in secondi,
// in corrispondenti giorni, ore, minuti, secondi

package main

import "fmt"

func main() {
	var sec uint
	const secondiGiorno = (3600 * 24)

	fmt.Print("Secondi?")
	fmt.Scan(&sec)

	fmt.Print("g:h:m:s -> ")

	fmt.Print(sec/secondiGiorno, ":")
	sec = sec % secondiGiorno
	fmt.Print(sec/3600, ":")
	sec = sec % 3600
	fmt.Print(sec/60, ":")
	sec = sec % 60
	fmt.Println(sec)
}
