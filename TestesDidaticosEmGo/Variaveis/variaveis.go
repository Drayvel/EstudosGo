package main

import "fmt"

func main() {
	//no Go há duas formas de declarar variáveis:
	//1 - Tipo Explicito
	//2 - Tipo Implicito
	var variavel1 string = "Variável 1" //variável explicita
	variavel2 := "Variável 2"           //variável implícita
	fmt.Println(variavel1)              //Imprime a mensagem que está armazenada na variável
	fmt.Println(variavel2)              //Imprime a mensagem que está armazenada na variável

	var (
		variavel3 string
		variavel4 string
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
