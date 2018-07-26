package main

import (
	"fmt"
	"strconv"
)

// Definindo tipo personalizado ou structs
type Carro struct {
	nome  string
	ano   int
	preco float32
	cor   string
}

func main() {
	/*
		Comentário 1
	*/

	// Definição de variável com tipagem estática
	var text_start string = "Inicío do olá mundo"
	var text_end string
	var count int = 4546846 * 3

	// Definição de variável com tipagem automática
	author := "Vagner Praia"
	print_author := true

	text_end = "Fim do olá mundo"

	// Definindo constantes
	const ano int = 2018

	fmt.Println("Texto inicial: " + text_start)
	fmt.Println(count)
	fmt.Println("Texto final: " + text_end)
	fmt.Println(print_author)
	fmt.Println("Author: " + author)
	fmt.Println("Ano: " + strconv.Itoa(ano))

	// Utilizando tipo personalizado
	var carro1 = Carro{
		nome:  "Teste Um",
		ano:   2018,
		preco: 50000.85,
		cor:   "Preto",
	}

	var carro2 = Carro{
		"Teste Dois",
		2018,
		30000.0,
		"Preto",
	}

	var carro3 = Carro{"Teste Três", 2017, 80000, "Preto"}

	fmt.Println("Nome: " + carro1.nome)
	fmt.Println(carro2)
	fmt.Println(carro3)

	mostrarNome(carro3)
}

func mostrarNome(carro Carro) {
	fmt.Println(carro.nome)
}
