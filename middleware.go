package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/caixingyi/rssagg/internal/auth"
	"github.com/caixingyi/rssagg/internal/database"
)

type authHandle func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, http.StatusUnauthorized, fmt.Sprintf("Error getting API key: %v", err))
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				responseWithError(w, http.StatusUnauthorized, "Invalid API key")
				return
			}
			responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
