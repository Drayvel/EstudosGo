package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("Função F")
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	//no Go não deixa você criar uma variável e não usar, por isso é necessário criar 2 variáveis para ter os dois resultados
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	//Se caso você não quiser um dos resultados, você pode usar o "_" pois assim o Go vai ignorar a segunda. Ele ainda vai dar dois retornos, mas ele só vai mostrar 1
	resultSoma, _ := calculosMatematicos(20, 30)
	fmt.Println(resultSoma)
}
