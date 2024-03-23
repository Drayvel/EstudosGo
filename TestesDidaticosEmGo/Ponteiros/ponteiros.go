package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10

	var variavel2 int = variavel1 //variavel2 copia a 1 neste momento
	fmt.Println(variavel1, variavel2)
	//A partir daqui a variavel 2 já não vai mais ser alterada
	variavel1++ //portanto apenas a 1 vai mudar
	fmt.Println(variavel1, variavel2)

	//Portanto é isso que acontece quando você atribui valores usando a cópia.

	//PONTEIROS:
	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro) //desreferenciação
	//ele está apontando pro endereço de memória, quando eu uso o asterisco eu digo pra ele que quero o valor dentro desse endereço,

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)

}
