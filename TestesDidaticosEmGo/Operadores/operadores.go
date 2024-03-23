package main

import "fmt"

func main() {
	//Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM DOS ARITMETICOS

	//ATRIBUIÇÃO

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//FIM DOS OPERADORES DE ATRIBUIÇÂO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	//FIM DOS OPERADORES RELACIONAIS

	//OPERADORES LÓGICOS
	fmt.Println(true && false) //retorna false
	fmt.Println(false || true) //retorna true
	fmt.Println(!true)         //retorna false
	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15        // Igual a numero + 15
	fmt.Println(numero) //imprime 11 pois o valor de numero é incrementado em 1

	numero--
	numero -= 20
	numero *= 3 //numero * 3
	numero /= 10
	numero %= 3

	fmt.Println(numero) //Imprime

	//FIM DOS OPERADORES UNARIOS

	//OPERADORES TERNARIOS
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
