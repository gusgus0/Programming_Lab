/*
Scrivere un programma che, data una retta, y = m*x + q, ed una parabola, y = a*x*x + b*x + c,
calcoli i punti di intersezione
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var m, q float64 = 2, 5       // y = 2x + 5
	var a, b, c float64 = 1, 2, 1 // y = 1x^2 + 2x + 1
	var x1, x2, y1, y2 float64

	fmt.Println("Retta: y = 2x + 5")
	fmt.Println("Parabola: y = 1x^2 + 2x + 1")

	//calcolo del delta
	var delta float64 = (b * b) + (m * m) - (2 * b * m) - 4*((a*c)-(a*q))

	if delta < 0 {
		fmt.Println("Delta negativo: La retta e la parabola non si intersecano fra loro")
	} else if delta == 0 {
		x1 = (-b + m) / (2 * a)
		y1 = (m * x1) + q
		fmt.Println("La retta e la parabola si intersecano nel punto P(", x1, y1, ")")
	} else {
		x1 = (-b + m + math.Sqrt(delta)) / (2 * a)
		x2 = (-b + m - math.Sqrt(delta)) / (2 * a)
		y1 = (m * x1) + q
		y2 = (m * x2) + q
		fmt.Println("La retta e la parabola si intersecano nei punti P(", x1, y1, ") e T(", x2, y2, ")")
	}
}
