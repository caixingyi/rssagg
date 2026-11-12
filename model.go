package main

import (
	"time"

	"github.com/caixingyi/rssagg/internal/database"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ApiKey    string `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID.String(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		ApiKey:    user.ApiKey,
	}
}
