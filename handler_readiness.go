package main

import "net/http"

func handlerReadliness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, struct{}{})
}
