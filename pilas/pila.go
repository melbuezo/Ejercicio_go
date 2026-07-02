package main

import "fmt"

type Pila struct {
	tope    *Nodo
	tamanio int
}

// Push
func (pila *Pila) Push(valor interface{}) {
	nuevoNodo := NuevoNodo(valor)
	nuevoNodo.Siguiente = pila.tope
	pila.tope = nuevoNodo
	pila.tamanio++
}

//pop
func (pila *Pila) Pop() (interface{}, error) {
	if pila.EsVacia() {
		return nil, fmt.Errorf("la pila está vacía")

	}
}
func (pila *Pila) EsVacia() bool {
	return pila.tope == nil
}
