package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&num1)

	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println("El número mayor es:", num1)
	} else {
		fmt.Println("El número mayor es:", num2)
	}
}
 