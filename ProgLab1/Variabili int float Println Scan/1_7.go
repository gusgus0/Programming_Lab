// Scrivere una applicazione Go che esprima una certa quantità di denaro (intero positivo) espressa in unità (ad es. Euro)
// nel numero più piccolo possibile di banconote/monete corrispondenti.
// Si supponga che i tagli disponibili siano: 100, 50, 20, 10, 5, 2, 1

package main

import "fmt"

func main() {
	var denaro int

	fmt.Print("Ammontare in euro? ")
	fmt.Scan(&denaro)

	fmt.Println(denaro/100, "banconote da 100")
	denaro = denaro % 100
	fmt.Println(denaro/50, "banconote da 50")
	denaro = denaro % 50
	fmt.Println(denaro/20, "banconote da 20")
	denaro = denaro % 20
	fmt.Println(denaro/10, "banconote da 10")
	denaro = denaro % 10
	fmt.Println(denaro/5, "banconote da 5")
	denaro = denaro % 5
	fmt.Println(denaro/2, "banconote da 2")
	denaro = denaro % 2
	fmt.Println(denaro, "banconote da 1")
}
