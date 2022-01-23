//Saber si un numero es primo
package main

import "fmt"

func main() {
	var numero, acum, lim int
	fmt.Print("Ingrese un numero: ")
	fmt.Scanf("%d", &numero)
	fmt.Println()
	fmt.Println("Numero ingresado:\t", numero)
	fmt.Print("Numero de divisores:\t")
	for div := 1; div <= numero; div++ {
		if numero%div == 0 {
			fmt.Print("[", div, "]")
			acum++
			lim++
			if lim == 10 {
				fmt.Printf("\n\t\t\t")
				lim = 0
			}
		}
	}
	fmt.Println()
	if acum == 2 {
		fmt.Println("Cantidad de divisores:\t", acum)
		fmt.Println("¿Es un numero primo?:\t Si")
	} else {
		fmt.Println("Cantidad de divisores:\t", acum)
		fmt.Println("¿Es un numero primo?:\t No")
	}
}
