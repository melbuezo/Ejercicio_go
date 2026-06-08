package main

import "fmt"

func main() {
	fmt.Println("BIENVENIDO A LA APLICACIÓN DE 6TO COMPUTACIOÓN")
	nombre, err := leerTexto("INGRESE SU NOMBRE")
	if err != nil {
		fmt.Print("ERROR AL LEER EL NOMBRE", err)
		return
	}
	edad, err := leerTexto("INGRESE SU EDAD")
	if err != nil {
		fmt.Print("ERROR AL LEER LA EDAD", err)
		return
	}

	fmt.Printf("HI¡ %s, tienes %s años de edad \n", nombre, edad)

	//-verificar si es mayor de edad

	if edad >= "18" {
		fmt.Printf("%d eres mayor de edad ", edad)
	} else {
		fmt.Printf("%d, eres menor de edad ", edad)
	}
	//--- estructura 2

	switch edad {
	case "18":
		fmt.Printf("%d, tu edad es ", edad)
	case "19":
		fmt.Printf("%d, tu edad es ", edad)
	default:
		fmt.Printf("%d, tu edad es ", edad)
	}

	//-- estructura 3 ciclos

	for i := 0; i < 5; i++ {
		fmt.Printf("Iteraación %d \n", i)
	}

	//-- estructura 4 ciclo while
	cont := 0
	for cont < 5 {
		fmt.Printf("contador: %d \n ", cont)
		cont++
	}
	//--- estructura 5 ciclo do while
	cont2 := 0
	for {
		fmt.Println(cont2)
		cont2++
		if cont2 > 5 {
			break
		}
	}

}
