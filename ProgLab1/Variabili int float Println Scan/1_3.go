// Scrivere una applicazione Go che, lette da input due frazioni (una frazione è rappresentata da una coppia di interi),
// calcoli la frazione somma

package main

import "fmt"

func main() {
	var n1, d1, n2, d2, nsum, dsum int
	fmt.Print("Inserire numeratore e denominatore della prima frazione: ")
	fmt.Scan(&n1, &d1)
	fmt.Print("Inserire numeratore e denominatore della seconda frazione: ")
	fmt.Scan(&n2, &d2)

	dsum = d1 * d2
	nsum = n1*d2 + n2*d1

	fmt.Println(n1, "/", d1, "+", n2, "/", d2, "=", nsum, "/", dsum)
}
