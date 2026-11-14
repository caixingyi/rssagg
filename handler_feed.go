package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/caixingyi/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	var p params
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      p.Name,
		Url:       p.URL,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed: %v", err))
		return
	}
	responseWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}
	responseWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
