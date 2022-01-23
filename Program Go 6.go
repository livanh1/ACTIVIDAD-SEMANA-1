package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("ESCRIBA SU NOMBRE: ")
	fmt.Scanln(&name)

	fmt.Print("¿CUAL ES SU EDAD?: ")
	fmt.Scanln(&age)

	saludar(name, age)
}

func saludar(nombre string, edad int) {
	fmt.Println("Hola ", nombre)
	fmt.Printf("Tienes %d años\n", edad)
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
}
