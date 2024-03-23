package main

import "fmt"

//Struct é uma coleção de campos, cada um tem um nome e um tipo.
type usuario struct { //Definindo a estrutura do usuário
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Dayvd"
	u.idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Dayvd", 22, enderecoExemplo} //devo usar chaves e não parenteses neste caso. Inicializando a struct com valores diretos nos campos
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 23}
	fmt.Println(usuario3)
}
