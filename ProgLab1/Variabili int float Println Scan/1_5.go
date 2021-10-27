// Scrivere una applicazione Go che calcoli la distanza tra due punti del piano Cartesiano.
// Si utilizzi la funzione Sqrt del package math (che accetta come argomento un float64)

package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by float64

	fmt.Print("(x,y) del primo punto: ")
	fmt.Scan(&ax, &ay)
	fmt.Print("(x,y) del secondo punto: ")
	fmt.Scan(&bx, &by)

	fmt.Println("Distanza tra i punti (", ax, ",", ay, ")", "e (", bx, ",", by, "):", math.Sqrt(((ax-bx)*(ax-bx))+((ay-by)*(ay-by))))
}
