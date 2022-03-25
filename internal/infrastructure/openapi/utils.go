package openapi

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(
	w http.ResponseWriter,
	code int,
	payload interface{},
) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(response))
}
