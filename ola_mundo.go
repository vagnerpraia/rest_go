package main

import (
	"fmt"
)

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

	text_end = "Fim do olá mundo"

	fmt.Println("Texto inicial: " + text_start)
	fmt.Println(count)
	fmt.Println("Texto final: " + text_end)
	fmt.Println("Author: " + author)

	// Comentário 2
}
