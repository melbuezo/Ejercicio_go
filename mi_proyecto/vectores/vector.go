package main

import "fmt"

func main() {
	//Crear un vector enteros
	var numeros []int
	fmt.Println(numeros)
	//iniciar el vector con valores
	exdeRodrigo := []string{"YHUN", "LISBETH", "ALEJANDRA", "Crisley", "Stephanie"}
	fmt.Println(exdeRodrigo)

	//greagar elementos con append
	nu1 := append(numeros, 10, 50, 75, 100)
	fmt.Println(nu1)

	//recorrer valores
	for i, valor := range exdeRodrigo {
		fmt.Printf("indice %d: %s\n", i, valor)
	}
}
