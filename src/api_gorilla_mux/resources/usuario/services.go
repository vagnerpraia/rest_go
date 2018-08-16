package usuario

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/treinamento_go/src/api_gorilla_mux/model"
	"github.com/treinamento_go/src/api_gorilla_mux/util"
)

var response model.Response

func getUsuarios(w http.ResponseWriter, r *http.Request) {
	data, err := getUsuariosModel()

	if err == nil {
		response.Code = http.StatusOK
		response.Message = "Usuário(s) retornado(s)."
		response.Data = data
	} else {
		util.HandlerError(&err, http.StatusInternalServerError, "Ocorreu um erro.")
	}

	util.EncodeResponseJson(w, response)
}

func getUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	data, err := getUsuarioModel(id)

	if err == nil {
		response.Code = http.StatusOK
		response.Message = "Usuário retornado."
		response.Data = data
	} else {
		util.HandlerError(&err, http.StatusInternalServerError, "Ocorreu um erro.")
	}

	util.EncodeResponseJson(w, response)
}

func postUsuario(w http.ResponseWriter, r *http.Request) {
	var usuarioRequest Usuario
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&usuarioRequest)
	defer r.Body.Close()

	if err == nil {
		data, err := insertUsuario(usuarioRequest)

		if err == nil {
			response.Code = http.StatusOK
			response.Message = "Usuário criado."
			response.Data = data
		} else {
			util.HandlerError(&err, http.StatusInternalServerError, "Ocorreu um erro.")
		}
	} else {
		util.HandlerError(&err, http.StatusInternalServerError, "Ocorreu um erro.")
	}

	util.EncodeResponseJson(w, response)
}

func login(w http.ResponseWriter, r *http.Request) {
	var usuarioRequest Usuario
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&usuarioRequest)
	defer r.Body.Close()

	if err == nil {
		var usuario Usuario

		usuario, err := loginModel(usuarioRequest)

		if err == nil {
			if usuario.Nome == "" {
				response.Code = http.StatusNoContent
			} else {
				response.Code = http.StatusOK
				response.Message = "Usuário retornado."
				response.Data = usuario
			}
		} else {
			util.HandlerError(&err, http.StatusInternalServerError, "Ocorreu um erro.")
		}
	} else {
		util.HandlerError(&err, http.StatusInternalServerError, "Ocorreu um erro.")
	}

	util.EncodeResponseJson(w, response)
}
