package sobre

type Sobre struct {
	Nome   string `json:"nome" bson:"nome"`
	Versao string `json:"versao" bson:"versao"`
	Autor  string `json:"autor" bson:"autor"`
	Email  string `json:"email" bson:"email"`
}

var sobre = Sobre{Nome: "API Rest Go com gorilla/mux", Versao: "0.0.1", Autor: "Vagner Praia", Email: "vagnerpraia@gmail.com"}
