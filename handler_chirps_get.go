package main

import "net/http"

func (cfg *apiConfig) HandlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not retrive chirps", err)
		return
	}
	chirps := make([]Chirp, len(dbChirps))
	for i, chirp := range dbChirps {
		chirps[i] = Chirp(chirp)
	}
	respondWithJSON(w, http.StatusOK, chirps)
}
