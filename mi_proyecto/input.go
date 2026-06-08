package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
//read compartido entre funciones 
var reader = bufio.NewReader(os.Stdin)

func leerTexto(promt string) (string, error){
	fmt.Print(promt)//mostrar el mensaje mano de main
	texto, err := reader.ReadString('\n')
	if err != nil{
		return "", err
	}
	return strings.TrimSpace(texto), nil
}