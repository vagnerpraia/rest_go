package main

type Usuario struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Usuarios []Usuario

var usuarios = Usuarios{
	Usuario{1, "Maria da Silva", "maria@mailtest.com", "123456"},
	Usuario{2, "João da Silva", "joao@mailtest.com", "456789"},
	Usuario{3, "Paulo da Silva", "paulo@mailtest.com", "789123"},
}
