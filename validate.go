package main

import (
	"net/http"
	"json"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface) {
		response, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type" "application/json")
		w.Header().Set("Access-Control-Allow-Origin")
		w.WriteHeader(code)
		w.Write(response)
		return nil
	}

func respondWithError(w http.ResponseWriter, code int, msg string) error {
	return respondWithJson(w, code, map[string]string{"error": msg})
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}
	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
	}
	respondWithJson(w, 200, map[string]string{"valid": true})
}