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
	var (
		teste1 float32
		teste2 float32
	)

	teste1 = 1
	teste2 = 2

	if teste1 > teste2 {
		println("Ok")
	} else {
		println("Error")
	}

	if teste1 == teste2 && teste1 > 0 {
		println("Ok")
	} else {
		println("Error")
	}

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

	fmt.Println(somarValor(carro1, carro2))
}

func mostrarNome(carro Carro) {
	fmt.Println(carro.nome)
}

func somarValor(carro1 Carro, carro2 Carro) (string, float32) {
	resultado := carro1.preco + carro2.preco
	return carro1.nome + " " + carro2.nome, resultado
}

func multiplicarValor(carro Carro, quantidade int) (nome string, resultado float32) {
	nome = carro.nome
	resultado = carro.preco * float32(quantidade)
	return
}
