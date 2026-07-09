package main

import "fmt"

type Pila struct {
	tope    *Nodo
	tamanio int
}

func PilaNueva() *Pila {
	return &Pila{
		tope:    nil,
		tamanio: 0,
	}
}

//PUSH
func (pila *Pila) push(valor interface{}) {
	nuevoNodo := NuevoNodo(valor)
	nuevoNodo.Siguiente = pila.tope
	pila.tope = nuevoNodo
	pila.tamanio++
}

//VACIO INFINITO
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

// ============================================================================
// MÉTODOS DE LOS EJERCICIOS (Añadir al final de pila.go)
// ============================================================================

// CODIGO 1: Contar elementos pares
func (pila *Pila) ContarPares() int {
	pilaAux := PilaNueva()
	contador := 0

	// Desapilamos para procesar
	for !pila.EstaVacia() {
		val, _ := pila.Pop()
		pilaAux.push(val)

		// Validamos que sea un entero antes de operar
		if num, ok := val.(int); ok {
			if num%2 == 0 {
				contador++
			}
		}
	}

	// Restauramos la pila original
	for !pilaAux.EstaVacia() {
		val, _ := pilaAux.Pop()
		pila.push(val)
	}

	return contador
}

// CODIGO 2: Sumar todos los elementos
func (pila *Pila) SumarElementos() int {
	pilaAux := PilaNueva()
	suma := 0

	for !pila.EstaVacia() {
		val, _ := pila.Pop()
		pilaAux.push(val)

		// Solo sumar si el tipo de dato es int
		if num, ok := val.(int); ok {
			suma += num
		}
	}

	// Restauramos la pila original
	for !pilaAux.EstaVacia() {
		val, _ := pilaAux.Pop()
		pila.push(val)
	}

	return suma
}

func (pila *Pila) EncontrarMaximo() (int, error) {
	if pila.EstaVacia() {
		return 0, fmt.Errorf("error: la pila está vacía")
	}

	pilaAux := PilaNueva()
	var maximo int
	primero := true

	for !pila.EstaVacia() {
		val, _ := pila.Pop()
		pilaAux.push(val)

		if num, ok := val.(int); ok {
			if primero {
				maximo = num
				primero = false
			} else if num > maximo {
				maximo = num
			}
		}
	}

	for !pilaAux.EstaVacia() {
		val, _ := pilaAux.Pop()
		pila.push(val)
	}

	return maximo, nil
}

func (pila *Pila) EstaOrdenadaAscendente() bool {
	if pila.Tamanio() <= 1 {
		return true
	}

	pilaAux := PilaNueva()
	ordenada := true

	valAnterior, _ := pila.Pop()
	pilaAux.push(valAnterior)

	numAnterior, okAnterior := valAnterior.(int)

	for !pila.EstaVacia() {
		valActual, _ := pila.Pop()
		pilaAux.push(valActual)

		numActual, okActual := valActual.(int)

		if okAnterior && okActual {
			if numActual > numAnterior {
				ordenada = false
			}
		}
		numAnterior = numActual
		okAnterior = okActual
	}

	for !pilaAux.EstaVacia() {
		val, _ := pilaAux.Pop()
		pila.push(val)
	}

	return ordenada
}