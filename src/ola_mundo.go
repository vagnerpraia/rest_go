package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	// Exemplos de slice/array dinâmico
	var lista4 []string

	lista4 = append(lista4, "item novo 1")
	lista4 = append(lista4, "item novo 2")
	lista4 = append(lista4, "item novo 3")
	lista4 = append(lista4, "item novo 4")
	lista4 = append(lista4, "item novo 5")

	fmt.Println(lista4)
	fmt.Println(lista4[0])
	fmt.Println(lista4[0:3])

	// Exemplos de condicionais
	flag1 := true
	flag2 := false
	idade := 18

	if flag1 {
		fmt.Println("1: Ok")
	} else {
		fmt.Println("1: Error")
	}

	if flag1 && flag2 {
		fmt.Println("2: Ok")
	} else {
		fmt.Println("2: Error")
	}

	if flag1 || flag2 {
		fmt.Println("3: Ok")
	} else {
		fmt.Println("3: Error")
	}

	if idade >= 65 {
		fmt.Println("Idoso")
	} else if idade >= 18 {
		fmt.Println("Adulto")
	} else if idade >= 12 {
		fmt.Println("Adolescente")
	} else if idade < 12 {
		fmt.Println("Criança")
	}

	// Recebendo parâmetros da linha de comando
	args := os.Args
	nome_usuario := args[1]
	idade_usuario, err_conv_idade_usuario := strconv.Atoi(args[2])

	fmt.Println("Olá " + nome_usuario)
	if err_conv_idade_usuario != nil {
		fmt.Println("Ocorreu um erro com o parâmetro idade_usuario")
		showError(err_conv_idade_usuario)
	} else if idade_usuario >= 18 {
		fmt.Println("Acesso permitido")
	} else {
		fmt.Println("Acesso negado")
	}

	// Exemplo de for
	produtos := []string{"Item 1", "Item 2", "Item 3"}
	for i := 0; i < len(produtos); i++ {
		fmt.Println("Produto: " + produtos[i])
	}

	// Exemplo de foreach
	for _, produto := range produtos {
		fmt.Println("Produto: " + produto)
	}

	// Exemplo de switch
	dia_semana := 5
	switch dia_semana {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Erro")
	}

	//Exemplo de leitura de arquivo
	arquivo, erro_leitura_arquivo := ioutil.ReadFile("commands.txt")
	showError(erro_leitura_arquivo)

	fmt.Println(string(arquivo))

	//Exemplo de escrita de arquivo, reescrevendo arquivo já existente e criando um novo arquivo caso arquivo não exista
	escrita1 := ioutil.WriteFile("Teste.txt", []byte("Teste 1"), 0777)
	showError(escrita1)

	//Exemplo de escrita de arquivo, adicionando dados em arquivo já existente
	arquivo_escrita, err_open_escrita := os.OpenFile("Teste.txt", os.O_APPEND, 0777)
	showError(err_open_escrita)
	escrita2, err_escrita_2 := arquivo_escrita.WriteString("\nTeste 2")
	showError(err_escrita_2)
	arquivo_escrita.Close()

	fmt.Println(string(escrita2))

	arquivo2, erro_leitura_arquivo2 := ioutil.ReadFile("Teste.txt")
	showError(erro_leitura_arquivo2)
	fmt.Println(string(arquivo2))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
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
