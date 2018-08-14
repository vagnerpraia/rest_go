package sobre

import (
	"encoding/json"
	"net/http"
)

func getSobre(w http.ResponseWriter, r *http.Request) {
	data := "API Rest Go com gorilla/mux"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	enc.Encode(data)
}
