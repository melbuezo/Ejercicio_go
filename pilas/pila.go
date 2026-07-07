package main

import "fmt"

type Pila struct {
	tope    *Nodo
	tamanio int
}
func PilaNueva() *Pila {
    return &Pila{
        tope : nil,
        tamanio: 0,
    }
}


// Push
func (pila *Pila) Push(valor interface{}) {
	nuevoNodo := NuevoNodo(valor)
	nuevoNodo.Siguiente = pila.tope
	pila.tope = nuevoNodo
	pila.tamanio++
}

func (pila *Pila) EstaVacia() bool {
	return pila.tope == nil 
}
// Pop elimina y retorna el elemento del tope de la pila
func (pila *Pila) Pop() (interface{}, error) {
    if pila.EstaVacia() {
        return nil, fmt.Errorf("error: pila vacía")
    }
    valor := pila.tope.Valor
    pila.tope = pila.tope.Siguiente
    pila.tamanio--
    return valor, nil
}

// Peek retorna el elemento del tope sin eliminarlo
func (pila *Pila) Peek() (interface{}, error) {
    if pila.EstaVacia() {
        return nil, fmt.Errorf("error: pila vacía")
    }
    return pila.tope.Valor, nil
}

// EstaVacia verifica si la pila está vacía
func (pila *Pila) Si_EstaVacia() bool {
    return pila.tope == nil
}
// Ejercicio 1: Separar pila en pares e impares
func SepararParesImpares(p *Pila) (*Pila, *Pila) {
	pares := PilaNueva()
	impares := PilaNueva()
	temp := PilaNueva()

	// Vaciamos la pila original guardando el orden en temp
	for !p.EstaVacia() {
		valor, _ := p.Pop()
		numero := valor.(int)
		if numero%2 == 0 {
			pares.Push(numero)
		} else {
			impares.Push(numero)
		}
		temp.Push(numero)
	}

	// Restauramos la pila original para no dejarla vacía
	for !temp.EstaVacia() {
		valor, _ := temp.Pop()
		p.Push(valor)
	}

	return pares, impares
}

// Tamaño retorna el número de elementos en la pila
func (pila *Pila) Tamanio() int {
    return pila.tamanio
}

// Imprimir muestra todos los elementos de la pila
func (pila *Pila) Imprimir() {
    if pila.EstaVacia() {
        fmt.Println("Pila vacía")
        return
    }
    actual := pila.tope
    fmt.Print("Tope -> ")
    for actual != nil {
        fmt.Printf("%v -> ", actual.Valor)
        actual = actual.Siguiente
    }
    fmt.Println("nil")
}

// Vaciar elimina todos los elementos de la pila
func (pila *Pila) Vaciar() {
    pila.tope = nil
    pila.tamanio = 0
}
