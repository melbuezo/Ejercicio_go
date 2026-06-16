package main

import "fmt"

func main() {
	lista := Listas{}

	lista.InsertarInicioFinal(10)
	lista.InsertarInicioFinal(20)
	lista.InsertarInicioFinal(5)
	lista.InsertarInicioFinal(50)

	fmt.Println("--- Ejercicio 1 ---")
	lista.imprimir()

	fmt.Printf("Longitud actual: %d\n", lista.Longitud())
	if ult, ok := lista.UltimoValor(); ok {
		fmt.Printf("Último valor: %d\n", ult)
	}

	fmt.Println("\n--- Ejercicio 2 ---")
	lista.InsertarDespuesDelPenultimo(99)
	lista.imprimir()
	fmt.Printf("Nueva longitud: %d\n", lista.Longitud())

	fmt.Println("\n--- Ejercicio 3 ---")
	fmt.Println("BUSCAR UN DATO 100:", lista.Buscar(100))

	lista.Eliminar(10)
	lista.imprimir()
}
