package auxiliar

import "fmt"

//Escrever registra uma mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}

//Se uma função começa com letra minuscula, ela só vai ser visivel dentro do pacote que ela está!
//Se ela começa com letra Maiúscula, então ela poderá ser exportada para outros pacotes.
