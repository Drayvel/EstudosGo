package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	var array1 [5]int = [5]int{2, 4, 6, 8, 10} // Declaração de um Array com tamanho fixo (n
	//Todos os dados dentro do ARRAY tem que ser do mesmo tipo
	fmt.Println(array1)         // Imprime o array completo
	fmt.Printf("%v\n", array1)  // Imprime o array completo usando %v para formatar a saída
	fmt.Printf("%+v\n", array1) // Imprime o array completo com detalhes dos valores
	array1[0] = 1
	fmt.Println(array1) // Modificação direta de um elemento do array

	array2 := [5]string{"Pos1", "Pos2", "Pos3", "Pos4", "Pos5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)                   // Imprime o array sem especificar o tamanho
	fmt.Println(len(array3), cap(array3)) // Mostra o comprimento e capacidade do array

	//SLICE

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8} //Declaracao de uma slice
	fmt.Println(slice)                     //Imprimindo a slice
	//o Slice NÃO É um array
	//ele não é limitado igual ao array. Portanto não precisa especificar o seu tamanho.
	//O slice é dinamico e pode crescer ou diminuir.

	fmt.Println(reflect.TypeOf(slice)) //TypeOf devolve o tipo de uma variável
	fmt.Println(reflect.TypeOf(array3))

	fmt.Println("-----------------------------------------")

	slice = append(slice, 18)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	fmt.Println("-----------------------------------------")

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Pos Alterada"
	fmt.Println(slice2) //o Slice aponta para um array, ele é uma fatia do array

	//Arraus Internos

	fmt.Println("-----------------------------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3)) //length e capacidade

	//A função make() serve para criar arrays e slices. A diferença entre eles é que os arrays tem tamanho fixo
	//A função make() aloca memória para criar um novo array ou slice
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3))

	slice4 := make([]float32, 5)
	slice4 = append(slice4, 12)
	fmt.Println(slice4)
	fmt.Println(len(slice4), cap(slice4))

}
