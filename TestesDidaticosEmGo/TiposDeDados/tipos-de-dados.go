package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := -1000000000000000000
	fmt.Println(numero)

	//uint - unsigned int
	var numero2 uint32 = 1000 //não pode ser negativo, é positivo por padrão
	fmt.Println(numero2)      //imprime 4294967284, po

	//alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Print(numero3)

	//BYTE = uints
	var numero4 byte = 123
	fmt.Println(numero4)

	//FIM NUMEROS INTEIROS

	//float32, float64 suportam numeros Reais, não apenas os inteiros.
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123456.67
	fmt.Println(numeroReal3)

	//FIM NUMEROS REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//FIM STRiNGS

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Ocorreu um erro")
	fmt.Println(erro) //nil é o valor padrão de um erro em GoLang
}
