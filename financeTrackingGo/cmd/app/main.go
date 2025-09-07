package main

import (
	"log"
	"os"

	"github.com/Pdhenrique/FinanceTracking/internal/db"
	"github.com/Pdhenrique/FinanceTracking/internal/http"
	"github.com/Pdhenrique/FinanceTracking/pkg/user"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST == "" {
		log.Fatal("env DB_HOST não definido")
	}

	DB_USER := os.Getenv("DB_USER")
	if DB_USER == "" {
		log.Fatal("env DB_USER não definido")
	}

	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		log.Fatal("env DB_PASSWORD não definido")
	}

	DB_PORT := os.Getenv("DB_PORT")
	if DB_PORT == "" {
		log.Fatal("env DB_PORT não definido")
	}

	DB_NAME := os.Getenv("DB_NAME")
	if DB_NAME == "" {
		log.Fatal("env DB_NAME não definido")
	}

	conn, err := db.Connect(DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	if err != nil {
		log.Fatal("ERROR conectando ao db:", err)
	}
	defer conn.Close()

	userStorage := db.NewUserStorage(conn)
	userService := user.NewService(userStorage)
	handler := http.NewUserHandler(userService)

	server := http.NewServer(handler, "8080")
	server.Start()
	defer server.Stop()

	select {}
}
