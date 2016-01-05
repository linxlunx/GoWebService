package controllers

import (
	"net/http"

	"../helpers"
)

func Index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"message": "Hai",
		}
		helpers.WriteResponse(resp, w, r)
	}
}
