package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartRouter(port string, router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server running on port %s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
