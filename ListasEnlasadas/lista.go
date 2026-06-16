package main

import (
	"fmt"
)

type Listas struct {
	cabeza *Nodo
}

func (l *Listas) InsertarInicioFinal(dato int) {
	nuevo := &Nodo{dato: dato}
	if l.cabeza == nil {
		l.cabeza = nuevo
		return
	}
	actual := l.cabeza
	for actual.siguiente != nil {
		actual = actual.siguiente
	}
	actual.siguiente = nuevo
}

func (l *Listas) Longitud() int {
	contador := 0
	actual := l.cabeza
	for actual != nil {
		contador++
		actual = actual.siguiente
	}
	return contador
}

func (l *Listas) UltimoValor() (int, bool) {
	if l.cabeza == nil {
		return 0, false
	}
	actual := l.cabeza
	for actual.siguiente != nil {
		actual = actual.siguiente
	}
	return actual.dato, true
}

func (l *Listas) InsertarDespuesDelPenultimo(dato int) {
	if l.cabeza == nil || l.cabeza.siguiente == nil {
		l.InsertarInicioFinal(dato)
		return
	}
	actual := l.cabeza
	for actual.siguiente != nil && actual.siguiente.siguiente != nil {
		actual = actual.siguiente
	}
	nuevo := &Nodo{dato: dato, siguiente: actual.siguiente}
	actual.siguiente = nuevo
}

func (l *Listas) Buscar(dato int) bool {
	actual := l.cabeza
	for actual != nil {
		if actual.dato == dato {
			return true
		}
		actual = actual.siguiente
	}
	return false
}

func (l *Listas) Eliminar(dato int) {
	if l.cabeza == nil {
		return
	}
	if l.cabeza.dato == dato {
		l.cabeza = l.cabeza.siguiente
		return
	}
	actual := l.cabeza
	for actual.siguiente != nil {
		if actual.siguiente.dato == dato {
			actual.siguiente = actual.siguiente.siguiente
			return
		}
		actual = actual.siguiente
	}
}

func (l *Listas) imprimir() {
	actual := l.cabeza
	for actual != nil {
		fmt.Printf("%d", actual.dato)
		if actual.siguiente != nil {
			fmt.Print(" -> ")
		}
		actual = actual.siguiente
	}
	fmt.Println()
}
