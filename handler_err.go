package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, http.StatusInternalServerError, "Server error")
}
