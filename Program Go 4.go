package main

import "fmt"

func main() {

	var numero1, numero2 int // Declarar variables de tipo entero

	fmt.Print("Número 1: ")
	fmt.Scanln(&numero1)

	fmt.Print("Número 2: ")
	fmt.Scanln(&numero2)

	// Sumar
	suma := numero1 + numero2

	// Imprimir resultado
	fmt.Printf("LA SUMATORIA DE %d + %d ES IGUAL A %d", numero1, numero2, suma)

}
