package main

import (
	"log"
	"net/http"
	"log"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig {
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", http.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/healthz", handleReadiness)
	mux.Handle("/metrics", apiCfg.handleMetrics)
	mux.Handle("/reset", apiCfg.HandlerReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	} 

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServerHTTP(w, r)
	})
}