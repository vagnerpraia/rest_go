package usuario

import (
	"encoding/json"
	"net/http"

	"github.com/treinamento_go/src/api_gorilla_mux/model"

	"github.com/gorilla/mux"
	"github.com/treinamento_go/src/api_gorilla_mux/util"
)

func getUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dataResponse := getUsuariosModel()

	enc := json.NewEncoder(w)
	enc.Encode(dataResponse)
}

func getUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dataResponse := getUsuarioModel(id)

	enc := json.NewEncoder(w)
	enc.Encode(dataResponse)
}

func postUsuario(w http.ResponseWriter, r *http.Request) {
	var usuarioRequest Usuario
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&usuarioRequest)
	util.ShowError(err)
	defer r.Body.Close()

	dataResponse := insertUsuario(usuarioRequest)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.Encode(dataResponse)
}

func login(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var usuarioRequest Usuario

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&usuarioRequest)
	defer r.Body.Close()

	if err == nil {
		var usuario Usuario

		for _, u := range usuarios {
			if u.Email == usuarioRequest.Email && u.Senha == usuarioRequest.Senha {
				usuario = u
				break
			}
		}

		if usuario.Nome == "" {
			response.Code = http.StatusNoContent
		} else {
			response.Code = http.StatusOK
			response.Message = "Usuário retornado."
			response.Data = usuario
		}
	} else {
		response.Code = http.StatusInternalServerError
		response.Message = "Ocorreu um erro."
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)

	enc := json.NewEncoder(w)
	enc.Encode(response)
}
