package main

import "fmt"

//contruir una estructura
type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	//crear una instancia de la estructura
	p1 := Persona{"Ana", 30}
	p2 := Persona{Nombre: "Luis", Edad: 18}

	//mostrar los datos de las personas
	fmt.Println(p1.Nombre, p1.Edad)
	fmt.Println(p2)
	//modificar los datos de la persona 1
	p1.Edad = 40
	fmt.Println(p1)

	per := []Persona{
		{"Carlos", 25},
		{"Maria", 22},
		{"Jorge", 28},
	}

	fmt.Println(per)
	per2 := append(per, Persona{"Yun", 21})

	//correr el slice
	for i, dato := range per2 {
		fmt.Printf("DATO Persona, NOMBRE: %s, EDAD: %d\n", i, dato.Nombre, dato.Edad)

	}
}
