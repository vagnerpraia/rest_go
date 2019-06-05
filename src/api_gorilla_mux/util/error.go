package util

import (
	"fmt"

	"github.com/vagnerpraia/treinamento_go/src/api_gorilla_mux/model"
)

func HandlerError(e *error, code int, message string) model.Response {
	fmt.Println(e)

	response := model.Response{Code: code, Message: message}

	return response
}
