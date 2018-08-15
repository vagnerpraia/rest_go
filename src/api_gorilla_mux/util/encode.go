package util

import (
	"encoding/json"
	"net/http"

	"github.com/treinamento_go/src/api_gorilla_mux/model"
)

func EncodeResponseJson(w http.ResponseWriter, response model.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)

	enc := json.NewEncoder(w)
	enc.Encode(response)
}
