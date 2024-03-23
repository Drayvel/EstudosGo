package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")
	fmt.Println("-------------")

	numero := -5

	if numero > 15 {
		fmt.Println("É Maior que 15")
	} else {
		fmt.Println("Não é maior que 15")
	}

	fmt.Println("-------------")

	if outroNumero := numero; outroNumero > 0 { //a variavel criada no ifinit deixa de existir quando  o IF termina.
		fmt.Println("Numero é maior que zero")
	} else if numero < -10 {
		fmt.Println("Numero é menor que -10")
	} else {
		fmt.Println("Numero está entre 0 e -10")
	}
}
