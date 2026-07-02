package main

import "fmt"

func main() {
	lista := Listas{}
	lista.InsertarAlFinal(5)
	lista.InsertarAlFinal(10)
	lista.InsertarAlFinal(15)
	lista.InsertarAlFinal(10) 
	lista.InsertarAlFinal(20)
	lista.InsertarAlFinal(30)
	lista.Imprimir() // Imprime: 5 10 15 20 30 null

	fmt.Println("Buscar el dato 100:", lista.Buscar(100))

	lista.Eliminar(5)
	fmt.Print("Despues de eliminar 5: ")
	lista.Imprimir()

	central, err := lista.ElementoCentral()
	if err == nil {
		fmt.Println("Elemento central:", central)
	}

	lista.EliminarDuplicados()
	fmt.Print("Eliminar duplicados: ")
	lista.Imprimir()

	lista.Invertir()
	fmt.Print("Lista invertida: ")
	lista.Imprimir()
}
