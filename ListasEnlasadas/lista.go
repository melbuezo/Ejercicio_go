package main

import (
	"errors"
	"fmt"
)

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

// imprimir la lista
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

func (Lis *Listas) Eliminar(valor int) {
	if Lis.Cabeza == nil { // no puede haber cenizar sino fuego
		return
	}
	if Lis.Cabeza.Valor == valor {
		Lis.Cabeza = Lis.Cabeza.Siguiente
		return
	}
	actual := Lis.Cabeza
	for actual.Siguiente != nil && actual.Siguiente.Valor != valor {
		actual = actual.Siguiente
	}

	if actual.Siguiente != nil {
		actual.Siguiente = actual.Siguiente.Siguiente
	}

}

// encontrar la longitud de la lista
func (Lis *Listas) Longitud() int {
	countador := 0
	actual := Lis.Cabeza
	for actual != nil {
		countador++
		actual = actual.Siguiente
	}
	return countador
}

// funcion 2 contrar
func (Lis *Listas) Ultimo() (int, error) {
	if Lis.Cabeza == nil {
		return 0,
			errors.New("Lista vacia")
	}
	actual := Lis.Cabeza
	for actual.Siguiente != nil {
		actual = actual.Siguiente
	}
	return actual.Valor, nil
}

// InsertarDespues coloca un nuevo nodo despues del primero que tenga buscado
func (Lis *Listas) InsertarDespues(buscado, nuevo int) {
	actual := Lis.Cabeza
	for actual != nil {
		if actual.Valor == buscado {
			nuevoNodo := &Nodo{Valor: nuevo}
			nuevoNodo.Siguiente = actual.Siguiente
			actual.Siguiente = nuevoNodo
			return
		}
		actual = actual.Siguiente
	}
}

// Invertir da vuelta la lista (metodo clasico con tres punteros)
func (Lis *Listas) Invertir() {
	var prev, siguiente *Nodo
	actual := Lis.Cabeza
	for actual != nil {
		siguiente = actual.Siguiente
		actual.Siguiente = prev
		prev = actual
		actual = siguiente
	}
	Lis.Cabeza = prev
}

func (Lis *Listas) ElementoCentral() (int, error) {
	if Lis.Cabeza == nil {
		return 0, errors.New("Lista vacia")
	}
	lento := Lis.Cabeza
	rapido := Lis.Cabeza
	for rapido != nil && rapido.Siguiente != nil {
		lento = lento.Siguiente
		rapido = rapido.Siguiente.Siguiente
	}
	return lento.Valor, nil
}

// EliminarDuplicados elimina valores repetidos (solo deja primera ocurrencia)
func (Lis *Listas) EliminarDuplicados() {
	if Lis.Cabeza == nil {
		return
	}
	vistos := make(map[int]bool)
	vistos[Lis.Cabeza.Valor] = true
	actual := Lis.Cabeza
	for actual.Siguiente != nil {
		if vistos[actual.Siguiente.Valor] {
			actual.Siguiente = actual.Siguiente.Siguiente
		} else {
			vistos[actual.Siguiente.Valor] = true
			actual = actual.Siguiente
		}
	}
}
