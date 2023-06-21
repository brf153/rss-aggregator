package main

import (
	"time"

	"github.com/brf153/rss-aggregator/internal/database"
	// "github.com/google/uuid"
)

type User struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"created_id"`
	UpdatedAt time.Time `json:"updated_id"`
	Name      string    `json:"name"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}
