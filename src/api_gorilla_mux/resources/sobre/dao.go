package sobre

func getSobreDao() (sobre Sobre) {
	sobre = Sobre{
		Nome:   "API Rest Go com gorilla/mux",
		Versao: "0.0.1",
		Autor:  "Vagner Praia",
		Email:  "vagnerpraia@gmail.com",
	}

	return
}
