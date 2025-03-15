package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("chirpID")
	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println(id, err)
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
	}
	chirp, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}
	respondWithJSON(w, http.StatusOK, Chirp(chirp))
}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}
	chirps := make([]Chirp, len(dbChirps))
	for i, chirp := range dbChirps {
		chirps[i] = Chirp(chirp)
	}
	respondWithJSON(w, http.StatusOK, chirps)
}
