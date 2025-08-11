package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/IrinaFosteeva/User_system_layered/handler"
	"github.com/IrinaFosteeva/User_system_layered/repository"
	"github.com/IrinaFosteeva/User_system_layered/service"
)

func mustEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func main() {
	dbHost := mustEnv("DB_HOST", "localhost")
	dbPort := mustEnv("DB_PORT", "5432")
	dbUser := mustEnv("DB_USER", "postgres")
	dbPass := mustEnv("DB_PASS", "postgres")
	dbName := mustEnv("DB_NAME", "postgres")
	port := mustEnv("PORT", "8080")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer dbpool.Close()

	// repository -> service -> handler
	repo := repository.NewPostgresUserRepo(dbpool)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	r := mux.NewRouter()
	h.RegisterRoutes(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		log.Printf("listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctxShut, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShut); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}
	log.Printf("server stopped")
}
