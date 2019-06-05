package sobre

import (
	"github.com/vagnerpraia/treinamento_go/src/api_gorilla_mux/model"
)

var GetIndexRoute = model.Route{"GET", "/", "getIndex", getSobreService}
var GetSobreRoute = model.Route{"GET", "/sobre", "getSobre", getSobreService}
