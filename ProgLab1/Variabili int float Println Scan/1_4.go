// Scrivere una applicazione Go che, date due "rette" (una retta è descritta da un coefficiente angolare m e da un
// termine noto q, di tipo float64) calcoli le coordinate del punto di intersezione

package main

import "fmt"

func main() {
	var m1, q1, m2, q2, x, y float64

	fmt.Print("Retta 1: inserire il coefficiente angolare seguito da un termine noto: ")
	fmt.Scan(&m1, &q1)
	fmt.Print("Retta 2: inserire il coefficiente angolare seguito da un termine noto: ")
	fmt.Scan(&m2, &q2)

	x = (q2 - q1) / (m1 - m2)
	y = (m1 * x) + q1

	fmt.Println("Punto di intersezione: (", x, ";", y, ")")
}
