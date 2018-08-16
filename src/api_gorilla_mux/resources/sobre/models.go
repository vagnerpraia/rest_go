package sobre

type Sobre struct {
	Nome   string `json:"nome" bson:"nome"`
	Versao string `json:"versao" bson:"versao"`
	Autor  string `json:"autor" bson:"autor"`
	Email  string `json:"email" bson:"email"`
}
