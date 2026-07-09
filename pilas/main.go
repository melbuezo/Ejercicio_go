package main

import "fmt"

func main() {
	pila := Pila{}

	fmt.Println("CREACION DE UNA PILA")
	pila.Imprimir()

	fmt.Println("\nAGREGAR DATOS (PUSH)")
	pila.push(19)
	pila.push(13)
	pila.push(9)
	pila.push(65)
	pila.push(666)
	pila.push(45)
	pila.push(97)
	pila.push(67)
	pila.push(0)
	pila.push(1)
	pila.Imprimir()

	fmt.Println("\nTAMAÑO DE LA PILA")
	tamanio := pila.Tamanio()
	fmt.Println("TAMAÑO: ", tamanio)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("\nEJERCICIO 001 ")
	pares := pila.ContarPares() // <- Llamada al método de pila.go
	fmt.Println("Cantidad de números pares:", pares)
	fmt.Print("Verificación de la pila original: ")
	pila.Imprimir()

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("\nEJERCICIO 002 ")
	suma := pila.SumarElementos() // <- Llamada al método de pila.go
	fmt.Println("Suma total de enteros:", suma)
	fmt.Print("Verificación de la pila original: ")
	pila.Imprimir()

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("\nEJERCICIO 003 ")
	maximo, err := pila.EncontrarMaximo()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El valor máximo es:", maximo)
	}
	fmt.Print("Verificación de la pila original: ")
	pila.Imprimir()

	fmt.Println("/////////////////////////")
	fmt.Println("EJERCICIO 004 ")

	fmt.Println("¿La pila actual está ordenada?:", pila.EstaOrdenadaAscendente())

	pilaOrdenada := Pila{}
	pilaOrdenada.push(1)
	pilaOrdenada.push(3)
	pilaOrdenada.push(5)
	pilaOrdenada.push(7)
	fmt.Print("Nueva pila de prueba: ")
	pilaOrdenada.Imprimir()
	fmt.Println("¿Esta nueva pila está ordenada?:", pilaOrdenada.EstaOrdenadaAscendente())
}