package main

import (
	"github.com/IrinaFosteeva/User_system_layered/internal/handler"
	"github.com/IrinaFosteeva/User_system_layered/internal/repository"
	"github.com/IrinaFosteeva/User_system_layered/internal/service"
	"log"

	"github.com/IrinaFosteeva/User_system_layered/config"
	"github.com/IrinaFosteeva/User_system_layered/db"
	"github.com/IrinaFosteeva/User_system_layered/server"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	pool, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	userRepo := repository.NewPostgresUserRepo(pool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := mux.NewRouter()
	userHandler.RegisterRoutes(r)

	server.StartRouter(cfg.Port, r)
}
