package main

import "fmt"

func main() {
	fmt.Println("MAPS")
	fmt.Println("-------------")
	//Dentro dos colchetes é o tipo das chaves. Fora dos colchetes é o tipo dos valores
	//Exemplo:
	userTest := map[string]int{
		"Joao":  1,
		"Maria": 2,
	}
	fmt.Println(userTest)

	//Podemos mudar o valor fora dos colchetes, exemplo:
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Rodrigues",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus B",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}
