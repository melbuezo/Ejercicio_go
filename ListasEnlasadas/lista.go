package main

import "fmt"

type Listas struct {
	Cabeza *Nodo // puntero al primer nodo
}

func (Lis *Listas) Insertar(valor int) {
	nuevo := &Nodo{Valor: valor} // crear un nuevo nodo con el valor dado
	nuevo.Siguiente = Lis.Cabeza // el siguiente del nuevo nodo apunta a la cabeza actual
	Lis.Cabeza = nuevo           // la cabeza ahora es el nuevo nodo

}

// insertar al final de la lista
func (Lis *Listas) InsertarAlFinal(valor int) {
	nuevo := &Nodo{Valor: valor}
	if Lis.Cabeza == nil {
		Lis.Cabeza = nuevo
		return
	}
	actualizar := Lis.Cabeza
	for actualizar.Siguiente != nil {
		actualizar = actualizar.Siguiente
	}
	actualizar.Siguiente = nuevo
}

//imprimir la lista
func (Lis *Listas) Imprimir() {
	ListaCompleta := Lis.Cabeza
	for ListaCompleta != nil {
		fmt.Printf("%d ", ListaCompleta.Valor)
		ListaCompleta = ListaCompleta.Siguiente
	}
	fmt.Println("null")
}

// buscar,vericar, analizar si un valor existe en la lista
func (Lis *Listas) Buscar(valor int) bool {
	ListaCompleta := Lis.Cabeza
	for ListaCompleta != nil {
		if ListaCompleta.Valor == valor {
			return true
		}
		ListaCompleta = ListaCompleta.Siguiente
	}
	return false
}
