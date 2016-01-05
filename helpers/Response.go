package helpers

import (
	"encoding/json"
	"net/http"
)

type T interface{}

func CheckMessage(status int, message string) string {
	StatusAll := StatusList()
	Temp := StatusAll.(map[int]interface{})

	if val, found := Temp[status]; found {
		return val.(string)
	} else {
		return message
	}
}

func ShowResponse(status int, message string, query interface{}) T {
	MessageResp := CheckMessage(status, message)

	resp := map[string]interface{}{
		"status":  status,
		"message": MessageResp,
		"result":  query,
	}

	return T(resp)
}

func WriteResponse(resp interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
