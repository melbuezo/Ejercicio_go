package main

import (
	"fmt"
)

func main() {

	pila := PilaNueva()
	fmt.Println("Creacion de una pila")

	pila.Imprimir()

	fmt.Print("Agregar datos (PUSH)")
	pila.Push(20)
	pila.Push(15)
	pila.Push(14)
	pila.Push(13)
	pila.Push(12)
	pila.Push(11)
	pila.Push(10)
	pila.Push(6)
	pila.Push(4)
	pila.Push(2)

	pila.Imprimir()

	fmt.Println("Tamaño de la pila")
	tamanio := pila.Tamanio()
	fmt.Println("tamaño:", tamanio)

	// Ejercicio 1: Separar en pares e impares
	fmt.Println("\nEjercicio 1: Separar en pares e impares")
	pares, impares := SepararParesImpares(pila)
	fmt.Print("Pares: ")
	pares.Imprimir()
	fmt.Print("Impares: ")
	impares.Imprimir()

	// Ejercicio 2: Invertir la pila
	fmt.Println("\nEjercicio 2: Invertir pila")
	fmt.Print("Original: ")
	pila.Imprimir()
	invertida := InvertirPila(pila)
	fmt.Print("Invertida: ")
	invertida.Imprimir()

	//ejercicio 3
	pila.Push(20)
	pila.Push(15)
	pila.Push(14)
	pila.Push(13)
	pila.Push(12)
	pila.Push(11)
	pila.Push(10)
	pila.Push(6)
	pila.Push(4)
	pila.Push(2)

	pila.Imprimir()

	fmt.Println("Tamanio de la pila")
	tamanio = pila.Tamanio()
	fmt.Println("tamaño:", tamanio)

}
