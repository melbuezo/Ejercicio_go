package main

import "fmt"

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
}
