package response

import (
	"encoding/json"
	"net/http"
)

type error struct {
	Error string
}

//setDefaultConf set default configuration for response data
func setDefaultConf(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//Error return error response
func Error(w http.ResponseWriter, msg string) {
	setDefaultConf(w)
	encoder := json.NewEncoder(w)
	encoder.Encode(error{Error: msg})
}

//Success return success response
func Success(w http.ResponseWriter, data interface{}) {
	setDefaultConf(w)

	res := make(map[string]interface{})
	res["data"] = data

	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}