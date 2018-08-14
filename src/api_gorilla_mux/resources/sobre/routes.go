package sobre

import (
	"github.com/treinamento_go/src/api_gorilla_mux/model"
)

var GetIndexRoute = model.Route{"GET", "/", "getIndex", getSobre}
var GetSobreRoute = model.Route{"GET", "/sobre", "getSobre", getSobre}
