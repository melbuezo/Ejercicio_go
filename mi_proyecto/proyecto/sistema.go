package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Estudiante struct {
	Nombre      string
	Carnet      string
	Asistencias int
	Faltas      int
}

var estudiantes []Estudiante

func main() {
	cargarDatos()

	for {
		fmt.Println("\n===== SISTEMA DE ASISTENCIA =====")
		fmt.Println("1. Registrar estudiante")
		fmt.Println("2. Mostrar estudiantes")
		fmt.Println("3. Pasar asistencia")
		fmt.Println("4. Ver reporte")
		fmt.Println("5. Guardar datos")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			registrarEstudiante()
		case 2:
			mostrarEstudiantes()
		case 3:
			pasarAsistencia()
		case 4:
			reporteAsistencia()
		case 5:
			guardarDatos()
		case 6:
			guardarDatos()
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}

func registrarEstudiante() {
	var nombre, carnet string

	fmt.Print("Ingrese nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese carnet: ")
	fmt.Scanln(&carnet)

	nuevo := Estudiante{
		Nombre:      nombre,
		Carnet:      carnet,
		Asistencias: 0,
		Faltas:      0,
	}

	estudiantes = append(estudiantes, nuevo)

	fmt.Println("Estudiante registrado correctamente")
}

func mostrarEstudiantes() {
	if len(estudiantes) == 0 {
		fmt.Println("No hay estudiantes registrados")
		return
	}

	fmt.Println("\n===== LISTA DE ESTUDIANTES =====")

	for i, e := range estudiantes {
		fmt.Printf("%d. %s | Carnet: %s\n", i+1, e.Nombre, e.Carnet)
	}
}

func pasarAsistencia() {
	if len(estudiantes) == 0 {
		fmt.Println("No hay estudiantes registrados")
		return
	}

	fmt.Println("\n===== PASAR ASISTENCIA =====")
	fmt.Println("Fecha:", time.Now().Format("02-01-2006"))

	for i := range estudiantes {
		var respuesta string

		fmt.Printf("%s asistió? (1/0): ", estudiantes[i].Nombre)
		fmt.Scanln(&respuesta)

		if respuesta == "1" || respuesta == "0" {
			estudiantes[i].Asistencias++
		} else {
			estudiantes[i].Faltas++
		}
	}

	fmt.Println("Asistencia registrada")
}

func reporteAsistencia() {
	if len(estudiantes) == 0 {
		fmt.Println("No hay estudiantes registrados")
		return
	}

	fmt.Println("\n===== REPORTE DE ASISTENCIA =====")

	for _, e := range estudiantes {
		total := e.Asistencias + e.Faltas
		porcentaje := 0.0

		if total > 0 {
			porcentaje = (float64(e.Asistencias) / float64(total)) * 100
		}

		fmt.Println("--------------------------------")
		fmt.Println("Nombre:", e.Nombre)
		fmt.Println("Carnet:", e.Carnet)
		fmt.Println("Asistencias:", e.Asistencias)
		fmt.Println("Faltas:", e.Faltas)
		fmt.Printf("Porcentaje: %.2f%%\n", porcentaje)
	}
}

func guardarDatos() {
	archivo, err := os.Create("estudiantes.json")

	if err != nil {
		fmt.Println("Error al guardar datos")
		return
	}

	defer archivo.Close()

	json.NewEncoder(archivo).Encode(estudiantes)

	fmt.Println("Datos guardados correctamente")
}

func cargarDatos() {
	archivo, err := os.Open("estudiantes.json")

	if err != nil {
		return
	}

	defer archivo.Close()

	json.NewDecoder(archivo).Decode(&estudiantes)
}
