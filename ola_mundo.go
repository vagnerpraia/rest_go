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

	fmt.Println(multiplicarValor(carro1, 3))
	fmt.Println(adicionarImpostos(carro1))

	mostrarNomes(1, carro1.nome, carro2.nome)

	// Exemplos de array
	var lista1 [10]string
	lista2 := [10]string{"1", "2"}

	lista1[0] = "item1"
	lista1[1] = "item2"

	fmt.Println(lista1)
	fmt.Println(lista2)
	fmt.Println(lista1[1])

	// Exemplos de array multidimensionais
	var lista3 [3][2]string
	lista3[0][0] = "Item A1"
	lista3[0][1] = "Item A2"
	lista3[1][0] = "Item B1"
	lista3[1][1] = "Item B2"
	lista3[2][0] = "Item C1"
	lista3[2][1] = "Item C2"

	fmt.Println(lista3)
	fmt.Println(lista3[0])
	fmt.Println(lista3[2][1])
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

func adicionarImpostos(carro Carro) (precoFinal float32) {
	calcularImposto := func() float32 {
		return carro.preco * 2
	}

	precoFinal = calcularImposto()

	return
}

func mostrarNomes(teste int, nomes ...string) {
	fmt.Println(teste)
	for _, nome := range nomes {
		fmt.Println(nome)
	}
}
