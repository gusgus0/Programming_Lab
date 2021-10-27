// Scrivere un programma che verifichi se, nel sistema utilizzato, una variabile di tipo int occupa 32 o 64 bits.
// Sarà visualizzato il messaggio: "32-bit encoding for int on the current architecture"
// oppure: "64-bit encoding for int on the current architecture"

package main

import "fmt"

func main() {
	var a int

	a = 2147483647 // 2^31 - 1

	if a+1 != 2147483648 {
		fmt.Println("32-bit encoding for int on the current architecture")
	} else {
		fmt.Println("64-bit encoding for int on the current architecture")
	}
}
