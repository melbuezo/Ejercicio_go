package main

import (
	"fmt"
)

// Definición de la estructura
type Libro struct {
	Titulo string
	Autor  string
	Anio   int
}

func main() {
	
	biblioteca := []Libro{
		{"El Quijote", "Miguel de Cervantes", 1605},
		{"Cien años de soledad", "Gabriel García Márquez", 2014},
		{"El laberinto de los espíritus", "Carlos Ruiz Zafón", 2016},
		{"La chica del tren", "Paula Hawkins", 2015},
		{"Rayuela", "Julio Cortázar", 1963},
		{"El capital en el siglo XXI", "Thomas Piketty", 2014},
		{"La traición de Roma", "Santiago Posteguillo", 2009},
		{"Cualquier día", "Juan José Millás", 2015},
		{"El código Da Vinci", "Dan Brown", 2003},
		{"La ciudad de los prodigios", "Eduardo Mendoza", 1986},
	}

	// 1. Mostrar todos los libros
	fmt.Println("--- Lista de todos los libros disponibles ---")
	for _, l := range biblioteca {
		fmt.Printf("- %s | %s | %d\n", l.Titulo, l.Autor, l.Anio)
	}

	// 2. Realizar la búsqueda (Reto: buscar libros entre 2014 y 2015)
	fmt.Println("\n--- Resultado de la búsqueda (2014 - 2015) ---")
	encontrado := false
	for _, l := range biblioteca {
		if l.Anio >= 2014 && l.Anio <= 2015 {
			fmt.Printf("Encontrado: %s (%d)\n", l.Titulo, l.Anio)
			encontrado = true
		}
	}

	if !encontrado {
		fmt.Println("No se encontraron libros en ese rango de años.")
	}
}