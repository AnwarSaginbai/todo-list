package main

import (
	"github.com/AnwarSaginbai/todo-list/internal/adapters/db"
	"github.com/AnwarSaginbai/todo-list/internal/adapters/rest"
	"github.com/AnwarSaginbai/todo-list/internal/application/api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load the config: %v\n", err)
	}
	storage, err := db.NewAdapter(os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect to the database: %v\n", err)
	}
	application := api.NewService(storage)
	httpServer := rest.NewAdapter(application)
	log.Fatal(httpServer.Run("4040"))
}
