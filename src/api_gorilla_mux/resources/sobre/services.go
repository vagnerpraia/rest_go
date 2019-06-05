package sobre

import (
	"net/http"

	"github.com/vagnerpraia/treinamento_go/src/api_gorilla_mux/model"
	"github.com/vagnerpraia/treinamento_go/src/api_gorilla_mux/util"
)

var response model.Response

func getSobreService(w http.ResponseWriter, r *http.Request) {
	data := getSobreDao()

	response.Code = http.StatusOK
	response.Message = "Informação sobre a aplicação retornada."
	response.Data = data

	util.EncodeResponseJson(w, response)
}
