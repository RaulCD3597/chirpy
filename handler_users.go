package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RaulCD3597/chirpy/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func NewUserFormDBUser(u database.User) User {
	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Email:     u.Email,
	}
}

func (cfg *apiConfig) handleUserCreation(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong", err)
		return
	}

	dbUser, err := cfg.db.CreateUser(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error while creating user", err)
		return
	}

	user := NewUserFormDBUser(dbUser)
	respondWithJSON(w, http.StatusCreated, user)
}
