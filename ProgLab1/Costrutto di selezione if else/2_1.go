// Scrivere un programma che confronti due frazioni in modo "esatto" (assumendo i denominatori di entrambe diversi da zero)

package main

import "fmt"

func main() {
	var n1, d1, n2, d2 float64

	fmt.Print("num e den fraz 1: ")
	fmt.Scan(&n1, &d1)
	fmt.Print("num e den fraz 2: ")
	fmt.Scan(&n2, &d2)

	if (n1 / d1) < (n2 / d2) {
		fmt.Println(n1, "/", d1, "<", n2, "/", d2)
	} else {
		fmt.Println(n2, "/", d2, "<", n1, "/", d1)
	}
}
